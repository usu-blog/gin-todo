package routes

import (
	"../db"

	"github.com/gin-gonic/gin"
)

// Index page
func Index(ctx *gin.Context) {
	todos := db.GetAll()
	ctx.HTML(200, "index.html", gin.H{
		"todos": todos,
	})
}
