package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Registercontroller(c *gin.Context) {
	username := c.PostForm("username")
	pwd := c.PostForm("password")
	c_pwd := c.PostForm("c_password")

	fmt.Println("username: %s   pwd: %s", username, pwd)
	if pwd == c_pwd {
		c.String(http.StatusOK, "successed")
	} else {
		c.String(http.StatusOK, "Failed")
	}
}
