package main

import (
	"log"
	"warehouse/internal/configs"
	"warehouse/internal/controller"
	"warehouse/internal/db"
	"warehouse/logger"
)

func main() {

	// Load configs
	if err := configs.ReadSettings(); err != nil {
		log.Fatalf("Ошибка чтения настроек: %s", err)
	}

	// Initializing logger
	if err := logger.Init(); err != nil {
		return
	}
	logger.Info.Println("Loggers initialized successfully!")

	// Connecting to db
	if err := db.ConnectDB(); err != nil {
		return
	}
	defer db.CloseDB()
	logger.Info.Println("Connection to database established successfully!")

	// Initializing db-migrations
	if err := db.InitMigrations(); err != nil {
		return
	}
	logger.Info.Println("Migrations initialized successfully!")

	// Running http-server
	if err := controller.RunServer(); err != nil {
		return
	}
}
