package repository

import (
	"finalMygram/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user model.User) (model.User, error)
	Delete(ID int) error
	FindAll() ([]model.User, error)
	FindById(ID int) (model.User, error)
	Update(user model.User) (model.User, error)
	Login(user model.User) (model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(user model.User) (model.User, error) {
	err := r.db.Create(&user).Error
	return user, err
}

func (r *userRepository) Delete(ID int) error {
	err := r.db.Delete(&model.User{}, ID).Error
	return err
}

func (r *userRepository) FindAll() ([]model.User, error) {
	var users []model.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) FindById(ID int) (model.User, error) {
	var user model.User
	err := r.db.Find(&user, ID).Error
	return user, err
}

func (r *userRepository) Update(user model.User) (model.User, error) {
	err := r.db.Save(&user).Error
	return user, err
}

func (r *userRepository) Login(user model.User) (model.User, error) {
	userRes := model.User{}
	err := r.db.First(&userRes, "email = ?", user.Email).Take(&userRes).Error
	return userRes, err
}
