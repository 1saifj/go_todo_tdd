package repository

import (
	"github.com/1saifj/go_todo_tdd/src/domain/model"
	"github.com/1saifj/go_todo_tdd/src/parameters"
)

type UserRepository interface {
	CreateUser(parameters.UserParameters) (*model.User, error)
	FindUserByID(id int) (*model.User, error)
	FindUserByEmail(email string) (*model.User, error)
	Login(parameters.LoginParameters) (uint, error)
}
