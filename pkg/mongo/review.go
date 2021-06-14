package mongo

import (
	"context"

	"github.com/codepd/tasmac"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// AddReview saves the given review in the repository
func (s *Store) AddReview(ctx context.Context, r tasmac.Review) error {
	if _, err := s.GetBeer(ctx, r.BeerID); err != nil {
		return err
	}
	col := s.client.Database(s.c.Database).Collection(s.c.ReviewCollection)

	if _, err := col.InsertOne(ctx, r); err != nil {
		return errors.Wrapf(err, "insert doc error")
	}
	return nil
}

func decodeReviewsFromCursor(ctx context.Context, cur *mongo.Cursor) ([]tasmac.Review, error) {
	l := make([]tasmac.Review, 0)
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var r tasmac.Review
		err := cur.Decode(&r)
		if err != nil {
			return nil, errors.WithMessage(err, "decode cursor")
		}
		l = append(l, r)
	}
	if err := cur.Err(); err != nil {
		return nil, errors.WithMessage(err, "cursor error")
	}
	return l, nil
}

// GetAll returns all reviews for a given beer
func (s *Store) GetAllReviews(ctx context.Context, beerID string) ([]tasmac.Review, error) {
	col := s.client.Database(s.c.Database).Collection(s.c.BeerCollection)
	filter := bson.D{
		{
			Key:   "beerId",
			Value: beerID,
		},
	}

	cur, err := col.Find(ctx, filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, tasmac.Errorf(tasmac.ENOTFOUND, "Beer not found.")
		}
		return nil, errors.Wrap(err, "find doc error")
	}
	return decodeReviewsFromCursor(ctx, cur)
}
