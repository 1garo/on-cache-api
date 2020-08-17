package models

import (
	"log"
	"net/http"
)

type LOGIN struct {
	USER     string `json:"user" binding:"required"`
	EMAIL 	 string `json:"email" binding:"required"`
	ID       int    `json:"id"`
}

var login []*LOGIN


func GetUsers() []*LOGIN {
	return login
}

func GetLogin() []*LOGIN {
	return login
}

func createSHA() {
	log.Printf("createSHA(): function called.")
}

func SetLogin(_login *LOGIN, id *int32) []*LOGIN {
	login = append(login, _login)
	log.Printf("ID: %d", *id)
	log.Printf("SETLOGIN: %v", login[*id])
	login[*id].ID = int(*id)
	return login
}

func CheckUser(user string, login []*LOGIN) (int, int) {
	var code, id int
	log.Println(len(login))
	for i := 0; i < len(login); i++ {
		log.Printf("checkId(): %v %d", login[i], i)
		log.Printf("login[i].ID: %d", login[i].ID)
		if tempUser := login[i].USER; tempUser == user{
			code = http.StatusOK
			id = login[i].ID
			break
		} else {
			code = http.StatusBadRequest
			id = -1
		}
	}
	return code, id
}

func CheckId(id int, login []*LOGIN) int {
	var code int
	log.Println(len(login))
	for i := 0; i < len(login); i++ {
		log.Printf("checkId(): %v %d", login[i], i)
		log.Printf("login[i].ID: %d", login[i].ID)
		if tempId := login[i].ID; tempId == id {
			code = http.StatusOK
			break
		} else {
			code = http.StatusBadRequest
		}
	}
	return code
}

