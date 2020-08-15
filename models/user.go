package models

import "log"

type LOGIN struct {
	USER     string `json:"user" binding:"required"`
	PASSWORD string `json:"password" binding:"required"`
	ID       int    `json:"id"`
}

var login []*LOGIN

func SetLogin(_login *LOGIN, id *int32) []*LOGIN {
	login = append(login, _login)
	log.Printf("ID: %d", *id)
	log.Printf("SETLOGIN: %v", login[*id])
	login[*id].ID = int(*id)
	return login
}

func GetUsers() []*LOGIN {
	return login
}
func GetLogin() []*LOGIN {
	return login
}
