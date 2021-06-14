package service

import (
	"context"
	"time"

	"github.com/codepd/tasmac"
	"github.com/google/uuid"
)

type rs struct {
	r tasmac.ReviewStore
}

func NewReviewService(r tasmac.ReviewStore) tasmac.ReviewService {
	return &rs{r: r}
}

func (s *rs) AddBeerReview(ctx context.Context, r tasmac.ReviewAdd) error {
	if err := s.r.AddReview(ctx, tasmac.Review{
		ID:        uuid.New().String(),
		BeerID:    r.BeerID,
		FirstName: r.FirstName,
		LastName:  r.LastName,
		Score:     r.Score,
		Text:      r.Text,
		Created:   time.Now(),
	}); err != nil {
		return err
	}
	return nil
}

func (s *rs) AddSampleReviews(ctx context.Context, r []tasmac.ReviewAdd) error {
	for _, rr := range r {
		if err := s.r.AddReview(ctx, tasmac.Review{
			ID:        uuid.New().String(),
			BeerID:    rr.BeerID,
			FirstName: rr.FirstName,
			LastName:  rr.LastName,
			Score:     rr.Score,
			Text:      rr.Text,
			Created:   time.Now(),
		}); err != nil {
			return err
		}
	}
	return nil
}

func (s *rs) GetBeerReviews(ctx context.Context, beerID string) ([]tasmac.Review, error) {
	return s.r.GetAllReviews(ctx, beerID)
}
