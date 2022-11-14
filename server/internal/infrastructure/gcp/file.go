package gcp

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/url"
	"path"
	"regexp"
	"strings"

	"cloud.google.com/go/storage"
	"github.com/google/uuid"
	"github.com/kennygrant/sanitize"
	"github.com/reearth/reearth-cms/server/internal/usecase/gateway"
	"github.com/reearth/reearth-cms/server/pkg/asset"
	"github.com/reearth/reearth-cms/server/pkg/file"
	"github.com/reearth/reearthx/log"
	"github.com/reearth/reearthx/rerror"
	"github.com/samber/lo"
	"google.golang.org/api/iterator"
)

const (
	gcsAssetBasePath string = "assets"
	fileSizeLimit    int64  = 10 * 1024 * 1024 * 1024 // 10GB
)

type fileRepo struct {
	bucketName   string
	base         *url.URL
	cacheControl string
	host         string
}

func NewFile(bucketName, base, cacheControl, host string) (gateway.File, error) {
	if bucketName == "" {
		return nil, errors.New("bucket name is empty")
	}

	var u *url.URL
	if base == "" {
		base = fmt.Sprintf("https://storage.googleapis.com/%s", bucketName)
	}

	var err error
	u, _ = url.Parse(base)
	if err != nil {
		return nil, errors.New("invalid base URL")
	}

	return &fileRepo{
		bucketName:   bucketName,
		base:         u,
		cacheControl: cacheControl,
		host:         host,
	}, nil
}

func (f *fileRepo) ReadAsset(ctx context.Context, u string, p string) (io.ReadCloser, error) {
	op := getGCSObjectPath(u, p)
	if op == "" {
		return nil, rerror.ErrNotFound
	}

	sn := sanitize.Path(op)
	if sn == "" {
		return nil, rerror.ErrNotFound
	}

	return f.read(ctx, sn)
}

func (f *fileRepo) GetAssetFiles(ctx context.Context, u string) ([]gateway.FileEntry, error) {
	p := getGCSObjectPath(u, "")
	b, err := f.bucket(ctx)
	if err != nil {
		return nil, rerror.ErrInternalBy(err)
	}

	it := b.Objects(ctx, &storage.Query{
		Prefix: p,
	})

	var fileEntries []gateway.FileEntry
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			return nil, rerror.ErrInternalBy(err)
		}

		fe := gateway.FileEntry{
			// /22/2232222233333/hoge/tileset.json -> hoge/tileset.json
			Name: strings.TrimPrefix(strings.TrimPrefix(attrs.Name, p), "/"),
			Size: attrs.Size,
		}
		fileEntries = append(fileEntries, fe)
	}

	if len(fileEntries) == 0 {
		return nil, gateway.ErrFileNotFound
	}

	return fileEntries, nil
}

func (f *fileRepo) UploadAsset(ctx context.Context, file *file.File) (string, error) {
	if file == nil {
		return "", gateway.ErrInvalidFile
	}
	if file.Size >= fileSizeLimit {
		return "", gateway.ErrFileTooLarge
	}

	uuid := newUUID()

	// あああああ.png
	p := getGCSObjectPath(uuid, file.Path)
	if p == "" {
		return "", gateway.ErrInvalidFile
	}

	// uuid1/uuid2/_____.png

	if err := f.upload(ctx, p, file.Content); err != nil {
		return "", err
	}
	return uuid, nil
}

func (f *fileRepo) DeleteAsset(ctx context.Context, u string, fn string) error {
	p := getGCSObjectPath(u, fn)
	if p == "" {
		return gateway.ErrInvalidFile
	}

	sn := sanitize.Path(p)
	if sn == "" {
		return gateway.ErrInvalidFile
	}
	return f.delete(ctx, sn)
}

func (f *fileRepo) GetURL(a *asset.Asset) string {
	return getURL(f.host, a.UUID(), a.FileName())
}

func (f *fileRepo) read(ctx context.Context, path string) (io.ReadCloser, error) {
	if path == "" {
		return nil, rerror.ErrNotFound
	}

	bucket, err := f.bucket(ctx)
	if err != nil {
		log.Errorf("gcs: read bucket err: %+v\n", err)
		return nil, rerror.ErrInternalBy(err)
	}

	reader, err := bucket.Object(path).NewReader(ctx)
	if err != nil {
		if errors.Is(err, storage.ErrObjectNotExist) {
			return nil, rerror.ErrNotFound
		}
		log.Errorf("gcs: read err: %+v\n", err)
		return nil, rerror.ErrInternalBy(err)
	}

	return reader, nil
}

func (f *fileRepo) upload(ctx context.Context, filename string, content io.Reader) error {
	if filename == "" {
		return gateway.ErrInvalidFile
	}

	bucket, err := f.bucket(ctx)
	if err != nil {
		log.Errorf("gcs: upload bucket err: %+v\n", err)
		return rerror.ErrInternalBy(err)
	}

	object := bucket.Object(filename)
	if err := object.Delete(ctx); err != nil && !errors.Is(err, storage.ErrObjectNotExist) {
		log.Errorf("gcs: upload delete err: %+v\n", err)
		return gateway.ErrFailedToUploadFile
	}

	writer := object.NewWriter(ctx)
	writer.ObjectAttrs.CacheControl = f.cacheControl

	if _, err := io.Copy(writer, content); err != nil {
		log.Errorf("gcs: upload err: %+v\n", err)
		return gateway.ErrFailedToUploadFile
	}

	if err := writer.Close(); err != nil {
		log.Errorf("gcs: upload close err: %+v\n", err)
		return gateway.ErrFailedToUploadFile
	}

	return nil
}

func (f *fileRepo) delete(ctx context.Context, filename string) error {
	if filename == "" {
		return gateway.ErrInvalidFile
	}

	bucket, err := f.bucket(ctx)
	if err != nil {
		log.Errorf("gcs: delete bucket err: %+v\n", err)
		return rerror.ErrInternalBy(err)
	}

	object := bucket.Object(filename)
	if err := object.Delete(ctx); err != nil {
		if errors.Is(err, storage.ErrObjectNotExist) {
			return nil
		}

		log.Errorf("gcs: delete err: %+v\n", err)
		return rerror.ErrInternalBy(err)
	}
	return nil
}

func getGCSObjectPath(uuid, objectPath string) string {
	if uuid == "" || !IsValidUUID(uuid) {
		return ""
	}

	ss := strings.Split(objectPath, "/")
	ss = lo.Map(ss, func(s string, _ int) string {
		return fileName(s)
	})

	return path.Join(gcsAssetBasePath, uuid[:2], uuid[2:], strings.Join(ss, "/"))
}

func (f *fileRepo) bucket(ctx context.Context) (*storage.BucketHandle, error) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	bucket := client.Bucket(f.bucketName)
	return bucket, nil
}

func newUUID() string {
	return uuid.New().String()
}

func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}

var re = regexp.MustCompile(`[^a-zA-Z0-9-_.]`)

func fileName(s string) string {
	ss := re.ReplaceAllString(s, "-")
	return strings.ToLower(ss)
}

func getURL(host, uuid, fName string) string {
	url, _ := url.JoinPath(host, gcsAssetBasePath, uuid[:2], uuid[2:], fileName(fName))
	return url
}
