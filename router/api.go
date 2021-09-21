package router

import (
	"github.com/labstack/echo/v4"

	"ehaba_backend_golang/handler"
)

type API struct {
	Echo        *echo.Echo
	UserHandler handler.UserHandler
}

func (api *API) SetUpRouter() {
	//route for User
	api.Echo.GET("/user/sign-in", api.UserHandler.SignIn)
	api.Echo.GET("/user/sign-up", api.UserHandler.SignUp)

	//route for HomePage

	//route for Notification

}
