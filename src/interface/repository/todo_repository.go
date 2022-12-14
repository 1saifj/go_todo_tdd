package repository

import (
	"github.com/1saifj/go_todo_tdd/src/domain/model"
	"github.com/1saifj/go_todo_tdd/src/parameters"
	"gorm.io/gorm"
)

type todoRepository struct {
	db *gorm.DB
}

type TodoRepository interface {
	FindAll(parameters.FilterTodoParameters) ([]*model.Todo, error)
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &todoRepository{db}
}

func (tr *todoRepository) FindAll(params parameters.FilterTodoParameters) ([]*model.Todo, error) {
	var todos []*model.Todo
	todo := tr.db.Model(&todos).Limit(params.GetLimit()).Offset(params.GetOffset()).Order(params.GetOrderBy())
	if params.Title != nil {
		todo = todo.Where("title LIKE ?", "%"+*params.Title+"%")
	}
	if params.IsComplete != nil {
		todo = todo.Where("is_complete = ?", *params.IsComplete)
	}
	err := todo.Find(&todos).Error
	if err != nil {
		return nil, err
	}
	return todos, nil

}
