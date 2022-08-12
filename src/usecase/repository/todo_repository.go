package repository

import (
	"github.com/1saifj/go_todo_tdd/src/domain/model"
	"github.com/1saifj/go_todo_tdd/src/parameters"
)

type TodoRepository interface {
	FindAll(parameters.FilterTodoParameters) ([]*model.Todo, error)
}
