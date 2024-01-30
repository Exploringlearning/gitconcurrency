package services

import (
	"encoding/json"
	//"fmt"
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
	ch := make(chan dto.Page)
	
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
	 if error!= nil {
        log.Fatalln(error)
     }

	go readPages(page)
	go writePages(ch)
	log.Print(page.Users)
	return page.Users, error
}


func readPages(page dto.Page) <-chan dto.Page {
	log.Print(" Inside read pages")
	ch := make(chan dto.Page)
        ch <- page
    return ch
}

func writePages(out <-chan dto.Page) (pages []dto.Page) {
	log.Print ("writePages")
    for i :=0 ; i< len(out) ; i++ {
		page := <- out
		log.Print(" channel read ", page)
		pages = append(pages, page)  
		
    }
    return pages

}



