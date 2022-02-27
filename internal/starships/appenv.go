package starships

import (
	"starships/internal/starships/models"

	"github.com/unrolled/render"
)

// AppEnv holds application configuration data
type AppEnv struct {
	Render        *render.Render
	Version       string
	Env           string
	Port          string
	StarshipStore models.StarshipStorage
}

// CreateContextForTestSetup initialises an application context struct
// for testing purposes
func CreateContextForTestSetup() AppEnv {
	testVersion := "0.0.1"
	appEnv := AppEnv{
		Render:        render.New(),
		Version:       testVersion,
		Env:           "LOCAL",
		Port:          "3001",
		StarshipStore: CreateStarshipService(CreateMockDataSet()),
	}
	return appEnv
}
