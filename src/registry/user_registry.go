package registry

import (
	"github.com/1saifj/go_todo_tdd/src/interface/controller"
	"github.com/1saifj/go_todo_tdd/src/interface/repository"
)

func (r Registery) NewUserController() controller.UserController {
	return controller.NewUserController(r.NewUserRepository())
}

func (r Registery) NewUserRepository() repository.UserRepository {
	return repository.NewUserRepository(r.db)
}
