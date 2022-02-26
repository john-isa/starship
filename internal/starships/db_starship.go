package starships

import (
	"starships/internal/starships/models"

	"github.com/palantir/stacktrace"
)

// UserService will hold the connection and key db info
type StarshipService struct {
	StarshipList  map[int]models.Starship
	MaxStarshipID int
}

// NewUserService creates a new Carer Service with the system's database connection
func NewStarshipService(list map[int]models.Starship, count int) *StarshipService {
	return &StarshipService{
		StarshipList:  list,
		MaxStarshipID: count,
	}
}

func (service *StarshipService) ListStarships(filterBy string, value string) ([]models.Starship, error) {
	shipList, err := service.ListAllStarships()
	list := []models.Starship{}
	switch filterBy {
	case "none":
		return shipList, err

	case "name":
		for v := 0; v < len(shipList); v++ {
			if shipList[v].Name == value {
				list = append(list, shipList[v])
			}
		}
	case "class":
		for v := 0; v < len(shipList); v++ {
			if shipList[v].Class == value {
				list = append(list, shipList[v])
			}
		}
	case "status":
		for v := 0; v < len(shipList); v++ {
			if shipList[v].Status == value {
				list = append(list, shipList[v])
			}
		}
	}
	return list, nil
}

func (service *StarshipService) ListAllStarships() ([]models.Starship, error) {
	// TBD
	return []models.Starship{}, nil
}

// GetUser returns a single JSON document
func (service *StarshipService) GetStarships(i int) (models.Starship, error) {
	user, ok := service.StarshipList[i]
	if !ok {
		return models.Starship{}, stacktrace.NewError("Failure trying to retrieve user")
	}
	return user, nil
}

// AddUser adds a User JSON document, returns the JSON document with the generated id
func (service *StarshipService) NewStarship(u models.Starship) (models.Starship, error) {
	service.MaxStarshipID = service.MaxStarshipID + 1
	u.ID = service.MaxStarshipID
	service.StarshipList[service.MaxStarshipID] = u
	return u, nil
}

// UpdateUser updates an existing user
func (service *StarshipService) UpdateStarship(u models.Starship) (models.Starship, error) {
	id := u.ID
	_, ok := service.StarshipList[id]
	if !ok {
		return u, stacktrace.NewError("Failure trying to update user")
	}
	service.StarshipList[id] = u
	return service.StarshipList[id], nil
}

// DeleteUser deletes a user
func (service *StarshipService) DeleteStarship(i int) error {
	_, ok := service.StarshipList[i]
	if !ok {
		return stacktrace.NewError("Failure trying to delete user")
	}
	delete(service.StarshipList, i)
	return nil
}

// CreateMockDataSet initialises a database for test purposes. It returns a list of User objects
// as well as the new max object count
func CreateMockDataSet() (map[int]models.Starship, int) {
	list := make(map[int]models.Starship)
	list[0] = models.Starship{
		ID:    0,
		Name:  "Lorikeet",
		Class: "Scout",
		Armament: []models.Armament{
			{
				ID:       0,
				ShipId:   0,
				Title:    "Extended Range Sensor Package",
				Quantity: "500",
			},
		},
		Crew:   25,
		Image:  "https://no.image.yet",
		Value:  350000000.00,
		Status: "ACTIVE",
	}
	list[1] = models.Starship{
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
	}
	return list, len(list) - 1
}
