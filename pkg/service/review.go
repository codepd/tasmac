package service

import "github.com/codepd/tasmac"

type rs struct {
	r tasmac.ReviewRespository
}

func NewReviewService(r tasmac.ReviewRespository) tasmac.ReviewService {
	return &rs{r: r}
}

func (s *rs) AddBeerReview(r tasmac.ReviewAdd) {
	_ = s.r.AddReview(r)
}

func (s *rs) AddSampleReviews(r []tasmac.ReviewAdd) {
	for _, rr := range r {
		_ = s.r.AddReview(rr) // error handling omitted for simplicity
	}
}

func (s *rs) GetBeerReviews(beerID string) []tasmac.Review {
	return s.r.GetAllReviews(beerID)
}
