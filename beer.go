package tasmac

import "time"

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
	AddBeer(...BeerAdd) error
	AddSampleBeers([]BeerAdd)
	GetBeer(string) (Beer, error)
	GetBeers() []Beer
}

type BeerRepository interface {
	// AddBeer saves a given beer to the repository.
	AddBeer(BeerAdd) error
	// GetAllBeers returns all beers saved in storage.
	GetAllBeers() []Beer
	// GetBeer returns the beer with given ID.
	GetBeer(string) (Beer, error)
}
