package main 

import(
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    // Simple group: v1
    v1 := router.Group("/v1")
    {
        v1.GET("/login", func(c *gin.Context){
        	c.String(http.StatusOK,"1")
        	})
        v1.GET("/submit", func(c *gin.Context){
        	c.String(http.StatusOK,"2")
        	})
        v1.GET("/read", func(c *gin.Context){
        	c.String(http.StatusOK,"3")
        	})
    }

    // Simple group: v2
    v2 := router.Group("/v2")
    {
        v2.POST("/login", nil)
        v2.POST("/submit", nil)
        v2.POST("/read", nil)
    }

    router.Run(":8080")
}