package service

import (
	"crud-gin/model"
	"crud-gin/repository"
)

func GetAllUsers() ([]model.User, error) {
	return repository.GetAllUsers()
}
