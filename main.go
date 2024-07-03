package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")
	//			router.StaticFS("/static", http.Dir("static"))

	router.GET("/love", func (ctx *gin.Context)  {
		name := ctx.Query("name")
		
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"name" : name,
		})
	}) 

	router.Run(":8000")
}