package main

import (
	"os"

	"github.com/PedroCost4/movies-backend/config"
	"github.com/PedroCost4/movies-backend/handler"
	"github.com/PedroCost4/movies-backend/server"
	"github.com/joho/godotenv"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	logger.Info("Starting the application")

	godotenv.Load(".env")

	err := config.Init()
	if err != nil {
		logger.Errorf("Unable to start the application: %v\n", err)
		os.Exit(1)
	}

	handler.InitHandlers()

	err = server.InitServer()
	if err != nil {
		logger.Errorf("Unable to start the server: %v\n", err)
		os.Exit(1)
	}

	logger.Info("Application started")
}
