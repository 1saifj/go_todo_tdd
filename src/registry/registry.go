package registry

import (
	"github.com/1saifj/go_todo_tdd/src/interface/controller"
	"gorm.io/gorm"
)

type Registery struct {
	db *gorm.DB
}

func NewRegistery(db *gorm.DB) *Registery {
	return &Registery{
		db: db,
	}
}

func (r Registery) NewAppController() controller.AppController {
	return controller.AppController{
		UserController: r.NewUserController(),
	}
}
