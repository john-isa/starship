package main

import (
	"os"
	"strings"

	"starships/internal/starships"

	"github.com/sirupsen/logrus"
	"github.com/unrolled/render"
)

func init() {
	if strings.ToUpper(os.Getenv("ENV")) == "LOCAL" {
		logrus.SetFormatter(&logrus.TextFormatter{})
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetFormatter(&logrus.JSONFormatter{})
		logrus.SetLevel(logrus.InfoLevel)
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
	logrus.WithFields(logrus.Fields{
		"env":     env,
		"path":    os.Getenv("VERSION"),
		"version": version,
	}).Info("Loaded VERSION file")

	// ===========================================================================
	// Initialise data storage
	// ===========================================================================
	starshipStore := starships.CreateStarshipService(starships.CreateMockDataSet())

	// ===========================================================================
	// Initialise application context
	// ===========================================================================
	appEnv := starships.AppEnv{
		Render:        render.New(),
		Version:       version,
		Env:           env,
		Port:          port,
		StarshipStore: starshipStore,
	}
	// ===========================================================================
	// Start application
	// ===========================================================================
	starships.StartServer(appEnv)
}
