package config

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/varadekd/card-game/api"
)

// SetupRouter is responsible for enabling HTTP requests using Gin for this application.
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// TODO: Uncomment the code when the application supports this mode.
	// If you wish to temporarily run this in release mode, you can use the command below.
	// RUN: export GIN_MODE=release

	// gin.SetMode(gin.ReleaseMode)

	// Creating server ping
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Calling all the apis
	api.SetupDeckApi(router)
	return router
}

// StartServer will initiate the server on the specified port.
// If the provided port is empty or if the router is nil, the application will be closed.
func StartServer(r *gin.Engine, port string) {
	if port == "" {
		log.Fatalln("Unable to start the server because of missing port.")
	}

	if r == nil {
		log.Fatalln("Unable to start the server because of missing router.")
	}

	err := r.Run(fmt.Sprintf(":%s", port))

	if err != nil {
		log.Fatalf("Unable to start the application on port %s. Encountered an error %s while attempting to start the application.", port, err.Error())
	}
}
