package utils

import "finalMygram/model"

func UserResponseFunc(user model.User) model.UserResponse {
	userResponse := model.UserResponse{}
	userResponse.ID = user.ID
	userResponse.Age = user.Age
	userResponse.Email = user.Email
	userResponse.Username = user.Username

	return userResponse
}
