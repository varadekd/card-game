package api

import (
	"github.com/gin-gonic/gin"
	"github.com/varadekd/card-game/controller"
)

func SetupDeckApi(r *gin.Engine) {
	r.POST("/deck/new", controller.GeneratedDeck)
	r.GET("/deck/:id", controller.OpenDeck)
}
