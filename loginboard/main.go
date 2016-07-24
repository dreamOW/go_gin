package main

import (
	"loginboard/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.Use(controller.Loginlimit, gin.Recovery())
	router.LoadHTMLGlob("view/*")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "temp1.html", gin.H{
			"title": "",
		})
	})
	router.GET("/toregister", func(c *gin.Context) {
		c.HTML(http.StatusOK, "temp2.html", gin.H{
			"title": "",
		})
	})
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "temp3.html", gin.H{
			"title": "",
		})
	})
	router.POST("/login", controller.Logincontroller)
	router.POST("/register", controller.Registercontroller)

	router.Run(":8080")
}
