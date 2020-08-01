package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"3chan/controllers"

	db "3chan/models"
)

func main() {
	db.Init()
	urls()
}

func urls() {
	route := gin.Default()
	route.Static("/static", "./views/static")
	route.LoadHTMLGlob("./views/templates/*.tmpl")
	route.GET("/threads/:threadId", controllers.GetThread)
	route.POST("/threads/:threadId", controllers.PostMessage)
	route.GET("/", controllers.GetThreadsList)
	route.POST("/", controllers.MakeThread)

	if err := route.Run(":8080"); err != nil {
		log.Fatal("Server Run Failed.: ", err)
	}
}
