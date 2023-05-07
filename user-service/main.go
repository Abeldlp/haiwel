package main

import (
	"github.com/Abeldlp/haiwel/user-service/config"
	"github.com/Abeldlp/haiwel/user-service/route"
	"log"
)

func main() {
	//Initialize env variables
	config.InitializeEnv()

	//Initialize database connection
	config.InitializeDatabase()
	defer config.CloseDatabase()

	//Initialize server
	server := config.InitializeHttpServer()
	route.InitializeRoutes()

	//Run http server
	err := server.Run(":8080")
	if err != nil {
		log.Fatalf("Could not listen on port 8080: %v", err)
	}
}
