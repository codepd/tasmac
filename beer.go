package tasmac

import (
	"context"
	"time"
)

type Beer struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Brewery   string    `json:"brewery"`
	Abv       float32   `json:"abv"`
	ShortDesc string    `json:"short_description"`
	Created   time.Time `json:"created"`
}

type BeerAdd struct {
	Name      string  `json:"name"`
	Brewery   string  `json:"brewery"`
	Abv       float32 `json:"abv"`
	ShortDesc string  `json:"short_description"`
}

// BeerService provides beer and review listing operations.
type BeerService interface {
	AddBeer(ctx context.Context, b BeerAdd) error
	AddSampleBeers(ctx context.Context, b []BeerAdd) error
	GetBeer(ctx context.Context, id string) (*Beer, error)
	GetBeers(ctx context.Context) ([]Beer, error)
}

type BeerStore interface {
	// AddBeer saves a given beer to the repository.
	AddBeer(ctx context.Context, b *Beer) error
	// GetAllBeers returns all beers saved in storage.
	GetAllBeers(ctx context.Context) ([]Beer, error)
	// GetBeer returns the beer with given ID.
	GetBeer(ctx context.Context, id string) (*Beer, error)
}
