package main

import (
	"log"
	"os"
	"tasklybe/pkg/server"
)

func main() {
	app := server.SetupApp()

	// Start server
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(app.Listen(":" + port))
}
