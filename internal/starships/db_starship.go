package starships

import (
	"errors"
	"fmt"
	"starships/internal/starships/models"
)

// StarshipService will hold the connection and key db info.
type StarshipService struct {
	StarshipList []models.Starship
}

// CreateStarshipService creates a new Starship Management Service with the system's database connection.
func CreateStarshipService(list []models.Starship) *StarshipService {
	return &StarshipService{
		StarshipList: list,
	}
}

func (service *StarshipService) ListAllStarships() ([]models.Starship, error) {
	shipList := service.StarshipList
	if len(shipList) == 0 {
		return shipList, errors.New("no starships in the inventory")
	}
	return shipList, nil
}

// ListStarships creates a list of sharships defined by a filter
func (service *StarshipService) ListStarships(filterBy string, value string) ([]models.Starship, error) {
	shipList, err := service.ListAllStarships()
	if err != nil {
		return []models.Starship{}, err
	}
	list := []models.Starship{}
	switch filterBy {
	case "none":
		return shipList, err

	case "name":
		for _, ship := range shipList {
			if ship.Name == value {
				list = append(list, ship)
			}
		}
	case "class":
		for _, ship := range shipList {
			if ship.Class == value {
				list = append(list, ship)
			}
		}
	case "status":
		for _, ship := range shipList {
			if ship.Status == value {
				list = append(list, ship)
			}
		}
	}
	if len(list) == 0 {
		str := fmt.Sprintf("no ship are available with the filter of type: %s and value %s", filterBy, value)

		return []models.Starship{}, errors.New(str)
	}
	return list, nil
}

// GetStarship gets the selected starship from the inventory
func (service *StarshipService) GetStarship(i int) (models.Starship, error) {
	ships := []models.Starship{}

	for _, ship := range service.StarshipList {
		if ship.ID == i {
			ships = append(ships, ship)
		}
	}

	if len(ships) == 0 {
		return models.Starship{}, errors.New("no ship entry found")
	}

	if len(ships) > 1 {
		return models.Starship{}, errors.New("error: ship ID entries")
	}

	return ships[0], nil
}

// CreateStarship creates and adds a starship to the inventory
func (service *StarshipService) CreateStarship(u models.Starship) (models.Starship, error) {
	service.StarshipList = append(service.StarshipList, u)

	return u, nil
}

// UpdateStarship updates an existing starship entry in the inventory
func (service *StarshipService) UpdateStarship(ship models.Starship) (models.Starship, error) {
	service.StarshipList[ship.ID] = ship

	return service.StarshipList[ship.ID], nil
}

// DeleteStarship deletes a starship from the inventory
func (service *StarshipService) DeleteStarship(i int) error {
	ship := service.StarshipList[i]

	service.StarshipList = append(service.StarshipList[:ship.ID], service.StarshipList[ship.ID+1])

	return nil
}

// CreateMockDataSet initialises a database for test purposes. It returns a list of User objects
// as well as the new max object count
func CreateMockDataSet() []models.Starship {
	list := []models.Starship{
		{
			ID:    0,
			Name:  "Lorikeet",
			Class: "Scout",
			Armament: []models.Armament{
				{
					ID:       0,
					ShipId:   0,
					Title:    "Extended Range Sensor Package",
					Quantity: "60",
				},
			},
			Crew:   25,
			Image:  "https://no.image.yet",
			Value:  350000000.00,
			Status: "ACTIVE",
		},
		{
			ID:    1,
			Name:  "Obliterator",
			Class: "Scout",
			Armament: []models.Armament{
				{
					ID:       0,
					ShipId:   1,
					Title:    "Ion Cannon",
					Quantity: "500",
				},
				{
					ID:       1,
					ShipId:   1,
					Title:    "Ion Cannon",
					Quantity: "500",
				},
			},
			Crew:   5472910,
			Image:  "https://no.image.yet",
			Value:  600000000000.00,
			Status: "ACTIVE",
		},
	}
	return list
}
