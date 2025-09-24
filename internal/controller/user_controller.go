package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ntva15z/go-ecommerce-backend-api/internal/service"
	"github.com/ntva15z/go-ecommerce-backend-api/pkg/response"
)

type UserController struct {
	userService service.IUserService
}

func NewUserController(
	userService service.IUserService,
) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) Register(c *gin.Context) {
	result := uc.userService.Register("", "")

	response.SuccessResponse(c, result, nil)
}
