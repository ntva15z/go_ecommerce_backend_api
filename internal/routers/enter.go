package routers

import (
	"github.com/ntva15z/go-ecommerce-backend-api/internal/routers/manage"
	"github.com/ntva15z/go-ecommerce-backend-api/internal/routers/user"
)

type RouterGroup struct {
	User   user.UserRouterGroup
	Manage manage.ManageRouterGroup
}

var RouterGroupApp = new(RouterGroup)
