package user

import (
	"github.com/gin-gonic/gin"
	"github.com/ntva15z/go-ecommerce-backend-api/internal/wire"
)

type UserRouter struct{}

func (u *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userController, _ := wire.InitUserRouterHandler()

	// public router
	uRouterPublic := Router.Group("/user")
	{
		uRouterPublic.GET("/register", userController.Register)
		uRouterPublic.POST("/opt")
	}

	// private router
	uRouterPrivate := Router.Group("/user")
	{
		uRouterPrivate.GET("/info")
		uRouterPrivate.POST("/opt")
	}
}
