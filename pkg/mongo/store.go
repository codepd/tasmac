package mongo

import (
	"context"
	"time"

	"github.com/codepd/tasmac/pkg/config"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Store struct {
	client *mongo.Client
	c      config.Mongo
}

func NewStore(conf config.Mongo) (*Store, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(conf.ConnectionURI))
	if err != nil {
		return nil, errors.WithMessage(err, "create mongo client")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := client.Connect(ctx); err != nil {
		return nil, errors.WithMessage(err, "mongo connect")
	}
	return &Store{
		client: client,
		c:      conf,
	}, nil
}
