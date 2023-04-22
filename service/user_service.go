package service

import (
	"finalMygram/model"
	"finalMygram/repository"
	"finalMygram/utils"
)

type UserService interface {
	Create(user model.User) (model.UserResponse, error)
	Delete(ID int) error
	FindAll() ([]model.UserResponse, error)
	FindById(ID int) (model.UserResponse, error)
	Update(ID int, newUser model.UserUpdate) (model.UserResponse, error)
	Login(user model.User) (model.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(r repository.UserRepository) *userService {
	return &userService{r}
}

func (s *userService) Create(user model.User) (model.UserResponse, error) {
	user, err := s.userRepository.Create(user)
	userResponse := utils.UserResponseFunc(user)
	return userResponse, err
}

func (s *userService) Delete(ID int) error {
	return s.userRepository.Delete(ID)
}

func (s *userService) FindAll() ([]model.UserResponse, error) {
	users, err := s.userRepository.FindAll()

	var usersResponse []model.UserResponse
	for _, user := range users {
		userResponse := utils.UserResponseFunc(user)
		usersResponse = append(usersResponse, userResponse)
	}

	return usersResponse, err
}

func (s *userService) FindById(ID int) (model.UserResponse, error) {
	user, err := s.userRepository.FindById(ID)
	userResponse := utils.UserResponseFunc(user)
	return userResponse, err
}

func (s *userService) Update(ID int, newUser model.UserUpdate) (model.UserResponse, error) {
	user, err := s.userRepository.FindById(ID)

	user.Username = newUser.Username
	user.Email = newUser.Email
	user.Age = newUser.Age

	newestUser, err := s.userRepository.Update(user)
	userResponse := utils.UserResponseFunc(newestUser)
	return userResponse, err
}

func (s *userService) Login(user model.User) (model.User, error) {
	return s.userRepository.Login(user)
}
