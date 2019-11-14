package main

import (
	"./db"
	"./routes"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	router := gin.Default()
	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/*.html")

	//Init()
	db.Init()

	//Index
	router.GET("/", routes.Index)

	router.Run()

}
