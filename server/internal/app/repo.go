package app

import (
	"context"
	"fmt"
	"time"

	"github.com/reearth/reearth-cms/server/internal/infrastructure/auth0"

	"github.com/reearth/reearth-cms/server/internal/infrastructure/fs"
	"github.com/reearth/reearth-cms/server/internal/infrastructure/gcs"
	"github.com/reearth/reearth-cms/server/pkg/log"
	"github.com/spf13/afero"

	mongorepo "github.com/reearth/reearth-cms/server/internal/infrastructure/mongo"
	"github.com/reearth/reearth-cms/server/internal/usecase/gateway"
	"github.com/reearth/reearth-cms/server/internal/usecase/repo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo"
)

func initReposAndGateways(ctx context.Context, conf *Config, debug bool) (*repo.Container, *gateway.Container) {
	repos := &repo.Container{}
	gateways := &gateway.Container{}

	// Mongo
	client, err := mongo.Connect(
		ctx,
		options.Client().
			ApplyURI(conf.DB).
			SetConnectTimeout(time.Second*10).
			SetMonitor(otelmongo.NewMonitor()),
	)
	if err != nil {
		log.Fatalf("repo initialization error: %+v\n", err)
	}
	if err := mongorepo.InitRepos(ctx, repos, client, ""); err != nil {
		log.Fatalf("Failed to init mongo: %+v\n", err)
	}

	// File
	datafs := afero.NewBasePathFs(afero.NewOsFs(), "data")
	var fileRepo gateway.File
	if conf.GCS.BucketName == "" {
		log.Infoln("file: local storage is used")
		fileRepo, err = fs.NewFile(datafs, conf.AssetBaseURL)
	} else {
		log.Infof("file: GCS storage is used: %s\n", conf.GCS.BucketName)
		fileRepo, err = gcs.NewFile(conf.GCS.BucketName, conf.AssetBaseURL, conf.GCS.PublicationCacheControl)
		if err != nil {
			if debug {
				log.Warnf("file: failed to init GCS storage: %s\n", err.Error())
				err = nil
			}
		}
	}
	if err != nil {
		log.Fatalln(fmt.Sprintf("file: init error: %+v", err))
	}
	gateways.File = fileRepo

	// Auth0
	gateways.Authenticator = auth0.New(conf.Auth0.Domain, conf.Auth0.ClientID, conf.Auth0.ClientSecret)

	return repos, gateways
}
