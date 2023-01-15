package decompressor

import (
	"archive/zip"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"cloud.google.com/go/storage"
	"github.com/bodgit/sevenzip"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/reearth/reearthx/log"
)

var (
	ErrUnsupportedExtention = errors.New("unsupoprted extention type")
)

const limit = 1024 * 1024 * 1024 * 30 // 30GB

const configPrefix = "REEARTH_CMS_WORKER"

type decompressor struct {
	zr  *zip.Reader
	sr  *sevenzip.Reader
	wFn func(name string) (io.WriteCloser, error)
}

func New(r io.ReaderAt, size int64, ext string, wFn func(name string) (io.WriteCloser, error)) (*decompressor, error) {
	if ext == "zip" {
		zr, err := zip.NewReader(r, size)
		if err != nil {
			return nil, err
		}
		return &decompressor{
			zr:  zr,
			wFn: wFn,
		}, nil
	} else if ext == "7z" {
		sr, err := sevenzip.NewReader(r, size)
		if err != nil {
			return nil, err
		}
		return &decompressor{
			sr:  sr,
			wFn: wFn,
		}, nil
	}
	return nil, ErrUnsupportedExtention
}

func (uz *decompressor) Decompress(assetBasePath string) error {
	zfs := []*zip.File{}
	if uz.zr != nil {
		for _, f := range uz.zr.File {
			fn := f.Name
			if strings.HasSuffix(fn, "/") {
				continue
			}
			if f.NonUTF8 {
				continue
			}
			if strings.HasPrefix(fn, "/") {
				continue
			}
			zfs = append(zfs, f)
		}
		uz.readConcurrentGCSFile(zfs, assetBasePath)
	} else if uz.sr != nil {
		for _, f := range uz.sr.File {
			if f.FileInfo().IsDir() {
				continue
			} else {
				rc, err := f.Open()
				if err != nil {
					return err
				}
				defer rc.Close()
				err = uz.read(f.Name, rc)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (uz *decompressor) read(name string, r io.Reader) error {
	w, err := uz.wFn(name)
	if err != nil {
		return err
	}
	_, err = io.CopyN(w, r, limit)
	_ = w.Close()
	if !errors.Is(err, io.EOF) && err != nil {
		for _, f := range uz.sr.File {
			return &LimitError{Path: f.FileInfo().Name()}
		}
	}
	return nil
}

func (uz *decompressor) readConcurrentGCSFile(zfs []*zip.File, assetBasePath string) {
	conf, cerr := ReadDecompressorConfig()
	if cerr != nil {
		log.Fatal(cerr)
	}
	var wg sync.WaitGroup  
	ctx := context.Background()
	client, _ := storage.NewClient(ctx)
	db := client.Bucket(conf.BucketName)
	workQueue := make(chan *zip.File, conf.DecompressorWorkerQueueDepth)
	for i := 0; i < int(conf.DecompressionWorkers); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			log.Infof("worker %d says hello!", i)
			for f := range workQueue {
				func() {
					fn := f.Name
					x, err := f.Open()
					if err != nil {
						log.Fatal(err)
					}
					defer x.Close()
					name := filepath.Join(assetBasePath, fn)
					w := db.Object(name).NewWriter(ctx)

					if _, err := io.Copy(w, x); err != nil {
						return
					}
					if err = w.Close(); err != nil {
						log.Infof("boom %s failed with %s", f.Name, err)
						return
					}
					log.Infof(" worker %d wrote %s!", i, f.Name)
				}()
			}
			log.Infof("Worker %d says bye!", i)
		}(i)
	}

	for _, f := range zfs {
		workQueue <- f
	}
	close(workQueue)
	wg.Wait()
}

type DecompressorConfig struct {
	DecompressionWorkers         int64  `envconfig:"DECOMPRESSION_NUM_WORKERS"`
	DecompressorWorkerQueueDepth int64  `envconfig:"DECOMPRESSION_WORKQUEUE_DEPTH"`	
	BucketName                   string `envconfig:"GCS_BUCKET_NAME"`
}

func ReadDecompressorConfig() (*DecompressorConfig, error) {
	if err := godotenv.Load(".env"); err != nil && !os.IsNotExist(err) {
		return nil, err
	} else if err == nil {
		log.Infof("config: .env loaded for decompressor")
	}

	var c DecompressorConfig
	err := envconfig.Process(configPrefix, &c)

	return &c, err
}

type LimitError struct {
	Path string
}

func (e *LimitError) Error() string {
	return fmt.Sprintf("file size limit reached at %s", e.Path)
}
