package routes

import (
	"github.com/gin-gonic/gin"

	"mtm-score-board/resources/models/response"
)

func ApplyRoutes(r *gin.Engine, router *Router) {
	r.GET("/profile", router.GetPlaythrough)
	r.GET("/score_board", router.ListPlaythrough)
	r.POST("/new_score", router.CreatePlaythrough)
}

func (r *Router) GetPlaythrough(c *gin.Context) {
	playerName := c.Query("name")
	playthroughHandler.GetPlaythrough(c, playerName)
}

func (r *Router) CreatePlaythrough(c *gin.Context) {
	var playthrough response.Record
	if c.BindJSON(&playthrough) != nil {
		c.Status(400)
		return
	}
	playthroughHandler.CreatePlaythrough(c, &playthrough)
}

func (r *Router) ListPlaythrough(c *gin.Context) {
	//If add auth later, it goes here.

	playthroughHandler.ListPlaythrough(c)
}
