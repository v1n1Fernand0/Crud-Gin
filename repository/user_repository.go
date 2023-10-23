package repository

import "crud-gin/model"

var users []model.User

func GetAllUsers() ([]model.User, error) {
	return users, nil
}
