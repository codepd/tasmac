package service

import (
	"errors"

	"github.com/codepd/tasmac"
)

// ErrDuplicate is used when a beer already exists.
var ErrDuplicate = errors.New("beer already exists")

type bs struct {
	r tasmac.BeerRepository
}

func NewBeerService(r tasmac.BeerRepository) tasmac.BeerService {
	return &bs{r: r}
}

func (s *bs) AddBeer(b ...tasmac.BeerAdd) error {
	// make sure we don't add any duplicates
	existingBeers := s.r.GetAllBeers()
	for _, bb := range b {
		for _, e := range existingBeers {
			if bb.Abv == e.Abv &&
				bb.Brewery == e.Brewery &&
				bb.Name == e.Name {
				return ErrDuplicate
			}
		}
	}

	// any other validation can be done here

	for _, beer := range b {
		_ = s.r.AddBeer(beer) // error handling omitted for simplicity
	}

	return nil
}

func (s *bs) AddSampleBeers(b []tasmac.BeerAdd) {
	// any validation can be done here
	for _, bb := range b {
		_ = s.r.AddBeer(bb) // error handling omitted for simplicity
	}
}

func (s *bs) GetBeer(id string) (tasmac.Beer, error) {
	return s.r.GetBeer(id)
}

func (s *bs) GetBeers() []tasmac.Beer {
	return s.r.GetAllBeers()
}
