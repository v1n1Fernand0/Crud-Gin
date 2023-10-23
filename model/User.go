package model

type User struct {
	Id int `json:"id"`
	Login string `json:"login"`
	Senha string `json:"senha"`	
}