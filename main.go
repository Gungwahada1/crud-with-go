package main

import (
	"fmt"
	"log"
	"net/http"

	// "fmt"
	"crud-with-go/config"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load variabel lingkungan dari file .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnectDB()

	messageServer := fmt.Sprintf("Server running on port %s", os.Getenv("SERVER_PORT"))
	hostServer := fmt.Sprintf("%s:%s", os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT"))

	log.Println(messageServer)
	http.ListenAndServe(
		hostServer,
		nil,
	)
}
