package user

import "github.com/gin-gonic/gin"

type ProductRouter struct{}

func (pr *ProductRouter) InitProductRouter(Router *gin.RouterGroup) {
	// public router
	prRouterPublic := Router.Group("/product")
	{
		prRouterPublic.GET("/search")
		prRouterPublic.GET("/detail/:id")
	}

	// private router
}
