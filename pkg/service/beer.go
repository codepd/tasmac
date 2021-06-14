package service

import (
	"context"
	"errors"
	"time"

	"github.com/codepd/tasmac"
	"github.com/google/uuid"
)

// ErrDuplicate is used when a beer already exists.
var ErrDuplicate = errors.New("beer already exists")

type bs struct {
	r tasmac.BeerStore
}

func NewBeerService(r tasmac.BeerStore) tasmac.BeerService {
	return &bs{r: r}
}

func (s *bs) AddBeer(ctx context.Context, b tasmac.BeerAdd) error {
	// make sure we don't add any duplicates
	existingBeers, err := s.r.GetAllBeers(ctx)
	if err != nil {
		return err
	}

	for _, e := range existingBeers {
		if b.Abv == e.Abv &&
			b.Brewery == e.Brewery &&
			b.Name == e.Name {
			return ErrDuplicate
		}
	}

	// any other validation can be done here

	if err := s.r.AddBeer(ctx, &tasmac.Beer{
		ID:        uuid.New().String(),
		Name:      b.Name,
		Brewery:   b.Brewery,
		Abv:       b.Abv,
		ShortDesc: b.ShortDesc,
		Created:   time.Now(),
	}); err != nil {
		return err
	}

	return nil
}

func (s *bs) AddSampleBeers(ctx context.Context, b []tasmac.BeerAdd) error {
	// any validation can be done here
	for _, beer := range b {
		if err := s.r.AddBeer(ctx, &tasmac.Beer{
			ID:        uuid.New().String(),
			Name:      beer.Name,
			Brewery:   beer.Brewery,
			Abv:       beer.Abv,
			ShortDesc: beer.ShortDesc,
			Created:   time.Now(),
		}); err != nil {
			return err
		}
	}
	return nil
}

func (s *bs) GetBeer(ctx context.Context, id string) (*tasmac.Beer, error) {
	return s.r.GetBeer(ctx, id)
}

func (s *bs) GetBeers(ctx context.Context) ([]tasmac.Beer, error) {
	return s.r.GetAllBeers(ctx)
}
