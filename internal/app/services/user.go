package services

import (
	"encoding/json"
	"ginconcurrency/internal/app/dto"
	"io"
	"log"
	"net/http"
)

type User interface{
	Get() ([]dto.User, error)
}
type user struct {

}

func NewUser() User {
    return &user{}
}

func (u *user) Get() ([]dto.User, error) {
	var page dto.Page 
	resp, err := http.Get("https://reqres.in/api/users?page=1")
	if err!= nil {
		log.Print("Not able to get the HTTP response ", err)
	}


	 body, err := io.ReadAll(resp.Body)
	 if err != nil {
		log.Fatalln(err)
	 }
     error := json.Unmarshal(body, &page)
	//  sb := string(body)
	 
	log.Print(page.Users)
	return page.Users, error
}

