package routes

import (
	"github.com/jjmoreno-dev/technical-test-interfell/controllers"
)

type UserRouteController struct {
	userController controllers.UserController
}

func NewRouteUserController(userController controllers.UserController) UserRouteController {
	return UserRouteController{userController}
}
