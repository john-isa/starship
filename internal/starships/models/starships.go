package models

import (
	"fmt"
)

// starships holds starships data
type Starship struct {
	ID       int        `json:"id"`
	Name     string     `json:"name"`
	Class    string     `json:"class"`
	Armament []Armament `json:"armament"`
	Crew     int        `json:"crew"`
	Image    string     `json:"image"`
	Value    float64    `json:"value"`
	Status   string     `json:"status"`
}

type Armament struct {
	ID       int    `json:"id"`
	ShipId   int    `json:"ship_id"`
	Title    string `json:"title"`
	Quantity string `json:"quantity"`
}

// StarshipStorage defines all the database operations
type StarshipStorage interface {
	ListStarships(filterBy string, value string) ([]Starship, error)
	ListAllStarships() ([]Starship, error)
	GetStarship(i int) (Starship, error)

	CreateStarship(u Starship) (Starship, error)
	UpdateStarship(u Starship) (Starship, error)
	DeleteStarship(i int) error
}

// GoString implements the GoStringer interface so we can display the full struct during debugging
// usage: fmt.Printf("%#v", i)
// ensure that i is a pointer, so might need to do &i in some cases
func (p *Starship) GoString() string {
	return fmt.Sprintf(`
{
	ID: %d,
	Name: %s,
	Class: %s,
	Crew: %d,
	Image: %s,
	Value: %f,
	Status: %s
}`,
		p.ID,
		p.Name,
		p.Class,
		p.Crew,
		p.Image,
		p.Value,
		p.Status,
	)
}
