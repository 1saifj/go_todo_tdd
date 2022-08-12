package router

import (
	"github.com/1saifj/go_todo_tdd/src/interface/controller"
	"github.com/kataras/iris/core/router"
)

func UserRouter(party router.Party, controller controller.UserController) {
	user := party.Party("/user")
	user.Post("", controller.CreateUser)
}
