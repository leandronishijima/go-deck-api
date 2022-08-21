package main

import "github.com/gin-gonic/gin"

var decks = []Deck{}

func setupRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.POST("/deck/new", CreateDeck)
		api.GET("/deck/open/:deck_id", OpenDeck)
		api.PATCH("/deck/:deck_id/draw", DrawCard)
	}

	return router
}

func main() {
	router := setupRouter()
	router.Run("localhost:8080")
}
