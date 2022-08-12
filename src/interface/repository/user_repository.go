package repository

import (
	"errors"

	"github.com/1saifj/go_todo_tdd/src/domain/model"
	"github.com/1saifj/go_todo_tdd/src/parameters"
	"github.com/1saifj/go_todo_tdd/src/usecase/repository"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{db}
}

func (ur *userRepository) CreateUser(params parameters.UserParameters) (*model.User, error) {
	user := model.User{
		Name:     params.FullName,
		Email:    params.Email,
		Password: params.Password,
	}
	if err := ur.db.Create(&user).Error; err != nil {
		return nil, err

	}
	return &user, nil
}

func (ur *userRepository) FindUserByEmail(email string) (*model.User, error) {
	var user model.User
	if err := ur.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
func (ur *userRepository) FindUserByID(id int) (*model.User, error) {
	var user model.User
	if err := ur.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *userRepository) Login(params parameters.LoginParameters) (uint, error) {
	user, err := ur.FindUserByEmail(params.Email)
	if err != nil {
		return 0, err
	}
	if user.Password != params.Password {
		return user.ID, nil
	}
	return 0, errors.New("invalid password")
}
