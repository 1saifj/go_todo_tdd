package controller

import (
	"github.com/1saifj/go_todo_tdd/src/interface/repository"
	"github.com/1saifj/go_todo_tdd/src/parameters"
	"github.com/kataras/iris"
)

type userController struct {
	UserRepository repository.UserRepository
}

type UserController interface {
	CreateUser(ctx iris.Context)
	Login(ctx iris.Context)
}

func NewUserController(ur repository.UserRepository) UserController {
	return &userController{UserRepository: ur}
}

func (uc userController) CreateUser(ctx iris.Context) {
	var params parameters.UserParameters
	err := ctx.ReadBody(&params)
	if err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}

	user, err := uc.UserRepository.CreateUser(params)
	if err != nil {
		return
	}

	ctx.StopWithJSON(iris.StatusCreated, user)

}

func (uc userController) Login(ctx iris.Context) {

}
