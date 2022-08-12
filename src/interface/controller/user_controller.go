package controller

import (
	"github.com/1saifj/go_todo_tdd/src/parameters"
	"github.com/1saifj/go_todo_tdd/src/usecase/repository"
	"github.com/kataras/iris/v12"
)

type userController struct {
	UserRepository repository.UserRepository
}

type UserController interface {
	CreateUser(ctx iris.Context)
	Login(ctx iris.Context)
}

func NewUserController(ur repository.UserRepository) {
	return &userController{ur}
}

func (uc userController) CreateUser(ctx iris.Context) {
	var params parameters.UserParameters
	err := ctx.ReadBody(&params)
	if err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}

	user, err := c.UserRepository.CreateUser(params)
	if err != nil {
		return
	}

	ctx.StopWithJSON(iris.StatusCreated, user)

}

func (uc userController) Login(ctx iris.Context) {

}
