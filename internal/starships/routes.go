package starships

// Route is the model for the router setup
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc HandlerFunc
}

// Routes are the main setup for our Router
type Routes []Route

var routes = Routes{
	Route{"Healthcheck", "GET", "/healthcheck", HealthcheckHandler},

	Route{"Getstarships", "GET", "/imperialship/{id}", ListAllStarshipsHandler},
	Route{"Getstarships", "GET", "/imperialship/{filterBy}/{value}", ListStarshipsHandler},
	Route{"Getstarships", "GET", "/imperialship/{id}", GetStarshipHandler},
	Route{"Createstarships", "POST", "/imperialship", NewStarshipHandler},
	Route{"Updatestarships", "PUT", "/imperialship/{pid:[0-9]+}", UpdateUserHandler},
	Route{"Deletestarships", "DELETE", "/imperialship/{pid:[0-9]+}", DeleteUserHandler},
}
