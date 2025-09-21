package initialize

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ntva15z/go-ecommerce-backend-api/internal/controller"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")

	{
		v1.GET("/ping", Pong)
		v1.GET("/user", controller.NewUserController().GetUserByID)
	}

	return r
}

func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
