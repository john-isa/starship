package starships

import (
	"encoding/json"
	"net/http"
	"strconv"

	models "starships/internal/starships/models"

	"starships/pkg/health"
	"starships/pkg/status"

	log "github.com/sirupsen/logrus"
)

// HandlerFunc is a custom implementation of the http.HandlerFunc
type HandlerFunc func(http.ResponseWriter, *http.Request, AppEnv)

// MakeHandler allows us to pass an environment struct to our handlers, without resorting to global
// variables. It accepts an environment (Env) struct and our own handler function. It returns
// a function of the type http.HandlerFunc so can be passed on to the HandlerFunc in main.go.
func MakeHandler(appEnv AppEnv, fn func(http.ResponseWriter, *http.Request, AppEnv)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// return function with AppEnv
		fn(w, r, appEnv)
	}
}

func HealthcheckHandler(w http.ResponseWriter, req *http.Request, appEnv AppEnv) {
	check := health.Check{
		AppName: "go-rest-api-template",
		Version: appEnv.Version,
	}
	appEnv.Render.JSON(w, http.StatusOK, check)
}

func ListAllStarshipsHandler(w http.ResponseWriter, req *http.Request, appEnv AppEnv) {
	//TBD
}

func ListStarshipsHandler(w http.ResponseWriter, req *http.Request, appEnv AppEnv) {
	//TBD
}

func GetStarshipHandler(w http.ResponseWriter, req *http.Request, appEnv AppEnv) {
	//TBD
}

func CreateStarshipHandler(w http.ResponseWriter, req *http.Request, appEnv AppEnv) {
	decoder := json.NewDecoder(req.Body)
	var u models.Starship
	err := decoder.Decode(&u)
	if err != nil {
		response := status.Response{
			Status:  strconv.Itoa(http.StatusBadRequest),
			Message: "malformed user object",
		}
		log.WithFields(log.Fields{
			"env":    appEnv.Env,
			"status": http.StatusBadRequest,
		}).Error("malformed user object")
		appEnv.Render.JSON(w, http.StatusBadRequest, response)
		return
	}
	starship := models.Starship{
		ID:       -1,
		Name:     u.Name,
		Class:    u.Class,
		Armament: u.Armament,
		Crew:     u.Crew,
		Image:    u.Image,
		Value:    u.Value,
		Status:   u.Status,
	}
	starship, _ = appEnv.StarshipStore.CreateStarship(starship)
	appEnv.Render.JSON(w, http.StatusCreated, starship)
}

// UpdateUserHandler updates a user object
func UpdateUserHandler(w http.ResponseWriter, req *http.Request, appEnv AppEnv) {
	// TBD
}

// DeleteUserHandler deletes a user
func DeleteUserHandler(w http.ResponseWriter, req *http.Request, appEnv AppEnv) {
	// TBD
}
