package manage

import "github.com/gin-gonic/gin"

type AdminRouter struct{}

func (ad *AdminRouter) InitAdminRouter(Router *gin.RouterGroup) {
	// private router
	uRouterPublic := Router.Group("/admin/user")
	{
		uRouterPublic.POST("/login")
	}

	// private router
	uRouterPrivate := Router.Group("/admin/user")
	{
		uRouterPrivate.POST("/active")
	}
}
