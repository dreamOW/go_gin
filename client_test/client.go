package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.GET("/setcookies.do", Printcontroller)
	}

	router.Run(":8089")
}

func Printcontroller(c *gin.Context) {
	ppinf, err := c.Cookie("ppinf")
	if err != nil {
		fmt.Println(err)
	}
	ppdig, err1 := c.Cookie("ppdig")
	if err1 != nil {
		fmt.Println(err1)
	}
	c.String(http.StatusOK, "hi : %s %s", ppinf, ppdig)
}
