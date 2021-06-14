package mongo

import (
	"context"

	"github.com/codepd/tasmac"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *Store) AddBeer(ctx context.Context, b *tasmac.Beer) error {
	col := s.client.Database(s.c.Database).Collection(s.c.BeerCollection)
	if _, err := col.InsertOne(ctx, b); err != nil {
		return errors.Wrapf(err, "insert doc error")
	}
	return nil
}

func decodeBeersFromCursor(ctx context.Context, cur *mongo.Cursor) ([]tasmac.Beer, error) {
	l := make([]tasmac.Beer, 0)
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var b tasmac.Beer
		err := cur.Decode(&b)
		if err != nil {
			return nil, errors.WithMessage(err, "cursor decode")
		}
		l = append(l, b)
	}
	if err := cur.Err(); err != nil {
		return nil, errors.WithMessage(err, "cursor error")
	}
	return l, nil
}

func (s *Store) GetAllBeers(ctx context.Context) ([]tasmac.Beer, error) {
	col := s.client.Database(s.c.Database).Collection(s.c.BeerCollection)
	cur, err := col.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	return decodeBeersFromCursor(ctx, cur)
}

func (s *Store) GetBeer(ctx context.Context, id string) (*tasmac.Beer, error) {
	col := s.client.Database(s.c.Database).Collection(s.c.BeerCollection)
	filter := bson.D{
		{
			Key:   "id",
			Value: id,
		},
	}
	var b tasmac.Beer
	if err := col.FindOne(ctx, filter).Decode(&b); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, tasmac.Errorf(tasmac.ENOTFOUND, "Beer not found.")
		}
		return nil, errors.Wrap(err, "find doc error")
	}
	return &b, nil
}
