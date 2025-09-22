package manage

import "github.com/gin-gonic/gin"

type UserRouter struct{}

func (u *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// private router
	uRouterPrivate := Router.Group("/admin/user")
	{
		uRouterPrivate.GET("/active")
	}
}
