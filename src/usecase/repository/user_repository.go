package repository

type UserRepository interface {
	CreateUser()
	Login()
	FindUserByID()
	FindUserByEmail()
}
