package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/varadekd/card-game/config"
	"github.com/varadekd/card-game/helper"
)

var router *gin.Engine

// TODO: Hardcoding values in the code is incorrect; it should retrieve values from the environment file.
const APP_PORT = "8080"

// The init function is responsible for enabling all the essential services
// required for this application's functionality.
func init() {
	fmt.Println("Setting up the initial services required by the server.")

	err := helper.GenerateDefaultDeck()

	if err != nil {
		log.Fatalln(err)
	}

	router = config.SetupRouter()
}

func main() {
	fmt.Printf("Starting application on port %s\n", APP_PORT)
	config.StartServer(router, APP_PORT)
}
