package tasmac

import "time"

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
	AddBeerReview(ReviewAdd)
	AddSampleReviews([]ReviewAdd)
	GetBeerReviews(string) []Review
}

type ReviewRespository interface {
	// GetAllReviews returns a list of all reviews for a given beer ID.
	GetAllReviews(string) []Review
	// AddReview saves a given review.
	AddReview(ReviewAdd) error
}
