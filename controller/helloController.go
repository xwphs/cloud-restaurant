package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HelloController struct {
}

// route
func (hello *HelloController) Router(engine *gin.Engine) {
	engine.GET("/hello", hello.HelloHandler)
	engine.DELETE("/user/:id", hello.DeleteHandler)
}

// hello handler
func (hello *HelloController) HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello cloud restaurant",
	})
}

// delete user handler
func (hello *HelloController) DeleteHandler(c *gin.Context) {
	id := c.Param("id")
	c.Writer.Write([]byte(fmt.Sprintf("delete user id: %v\n", id)))
}
