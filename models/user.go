package models

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
)

type LOGIN struct {
	USER  string `json:"user" binding:"required"`
	EMAIL string `json:"email" binding:"required"`
	ID    int    `json:"id"`
	HASH  string `json:"hash"`
}

var login []*LOGIN
var ok bool = false

func GetUsers() []*LOGIN {
	return login
}

func GetLogin() []*LOGIN {
	return login
}

func createSHA(user string, email string) string {
	log.Printf("createSHA(): function called.")
	msg := fmt.Sprintf("%s@%s", user, email)
	sha := sha1.New()
	sha.Write([]byte(msg))
	sha1_hash := hex.EncodeToString(sha.Sum(nil))
	log.Printf("SHA HASH: %s", sha1_hash)
	return sha1_hash
}

func SetLogin(_login *LOGIN, id *int32) []*LOGIN {
	login = append(login, _login)
	log.Printf("ID: %d", *id)
	log.Printf("SETLOGIN: %v", login[*id])
	userSha := createSHA(login[*id].USER, login[*id].EMAIL)
	login[*id].ID = int(*id)
	login[*id].HASH = userSha
	return login
}

// TODO: re-write check function to a unique generic check func
func CheckUserWithSHA(hash string, login []*LOGIN) (int, int, bool) {
	var code, id int
	log.Println(len(login))
	for i := 0; i < len(login); i++ {
		log.Printf("checkId(): %v %d", login[i], i)
		if tempHash := login[i].HASH; tempHash == hash {
			code = http.StatusOK
			ok = true
			break
		} else {
			code = http.StatusBadRequest
		}
	}
	return code, id, ok
}

func CheckId(id int, login []*LOGIN) (int, bool) {
	var code int
	log.Println(len(login))
	for i := 0; i < len(login); i++ {
		log.Printf("checkId(): %v %d", login[i], i)
		log.Printf("login[i].ID: %d", login[i].ID)
		if tempId := login[i].ID; tempId == id {
			code = http.StatusOK
			ok = true
			break
		} else {
			code = http.StatusBadRequest
		}
	}
	return code, ok
}
