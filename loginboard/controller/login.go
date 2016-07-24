package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manucorporat/stats"
)

var ips = stats.New()

func Sendcontroller(c *gin.Context) {
	c.String(http.StatusOK, "hi guy!!!!!!!!")
}

func Loginlimit(c *gin.Context) {
	ip := c.ClientIP()
	value := int(ips.Add(ip, 1))
	if value%50 == 0 {
		fmt.Printf("ip: %s, count: %d\n", ip, value)
	}
	if value >= 200 {
		if value%200 == 0 {
			fmt.Println("ip blocked")
		}
		c.Abort()
		c.String(503, "you were automatically banned :)")
	}
}

func Logincontroller(c *gin.Context) {
	username := c.PostForm("username")
	pwd := c.PostForm("password")

	fmt.Println("username: %s   pwd: %s", username, pwd)
	if username == "wangzhi" {
		if pwd == "12345" {
			c.Redirect(301, "/")
		} else {
			c.Redirect(301, "/index")
		}
	}
}
