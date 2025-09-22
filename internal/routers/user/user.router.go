package user

import "github.com/gin-gonic/gin"

type UserRouter struct{}

func (u *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// public router
	uRouterPublic := Router.Group("/user")
	{
		uRouterPublic.GET("/register")
		uRouterPublic.POST("/opt")
	}

	// private router
	uRouterPrivate := Router.Group("/user")
	{
		uRouterPrivate.GET("/info")
		uRouterPrivate.POST("/opt")
	}
}
