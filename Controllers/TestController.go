package Controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Hello in TestController
func Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello",
	})
}
