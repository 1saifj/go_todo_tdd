package router

import (
	"github.com/1saifj/go_todo_tdd/src/interface/controller"
	"github.com/kataras/iris/core/router"
)

func Routers(apiBuilder *router.APIBuilder, controller controller.UserController) {

	api := apiBuilder.Party("/api/v1")
	UserRouter(api, controller)

}
