package main

import (
	"os"
	"strings"

	starships "starships/internal/starships"

	log "github.com/sirupsen/logrus"
	"github.com/unrolled/render"
)

func init() {
	if strings.ToUpper(os.Getenv("ENV")) == "LOCAL" {
		log.SetFormatter(&log.TextFormatter{})
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetFormatter(&log.JSONFormatter{})
		log.SetLevel(log.InfoLevel)
	}
}

func main() {
	// ===========================================================================
	// Load environment variables
	// ===========================================================================
	var (
		env     = "LOCAL" // LOCAL, DEV, STG, PRD
		port    = "3001"  // server traffic on this port
		version = "0.1.0" // path to VERSION file
	)
	// ===========================================================================
	// Read version information
	// ===========================================================================
	// version, err := vparse.ParseVersionFile("VERSION")
	// if err != nil {
	// 	log.WithFields(log.Fields{
	// 		"env":  env,
	// 		"err":  err,
	// 		"path": os.Getenv("VERSION"),
	// 	}).Fatal("Can't find a VERSION file")
	// 	return
	// }
	log.WithFields(log.Fields{
		"env":     env,
		"path":    os.Getenv("VERSION"),
		"version": version,
	}).Info("Loaded VERSION file")
	// ===========================================================================
	// Initialise data storage
	// ===========================================================================
	userStore := starships.NewStarshipService(starships.CreateMockDataSet())
	// ===========================================================================
	// Initialise application context
	// ===========================================================================
	appEnv := starships.AppEnv{
		Render:    render.New(),
		Version:   version,
		Env:       env,
		Port:      port,
		UserStore: userStore,
	}
	// ===========================================================================
	// Start application
	// ===========================================================================
	starships.StartServer(appEnv)
}
