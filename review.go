package tasmac

import (
	"context"
	"time"
)

type Review struct {
	ID        string    `json:"id"`
	BeerID    string    `json:"beer_id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Score     int       `json:"score"`
	Text      string    `json:"text"`
	Created   time.Time `json:"created"`
}

type ReviewAdd struct {
	BeerID    string `json:"beer_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Score     int    `json:"score"`
	Text      string `json:"text"`
}

// ReviewService provides reviewing operations.
type ReviewService interface {
	AddBeerReview(ctx context.Context, r ReviewAdd) error
	AddSampleReviews(ctx context.Context, r []ReviewAdd) error
	GetBeerReviews(ctx context.Context, beerID string) ([]Review, error)
}

type ReviewStore interface {
	// AddReview saves a given review.
	AddReview(ctx context.Context, r Review) error
	// GetAllReviews returns a list of all reviews for a given beer ID.
	GetAllReviews(ctx context.Context, beerID string) ([]Review, error)
}
