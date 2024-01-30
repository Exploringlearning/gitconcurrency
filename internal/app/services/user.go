package services

import (
	"encoding/json"
	"time"

	"sync"
	//"fmt"
	"ginconcurrency/internal/app/dto"
	"io"
	"log"
	"net/http"
)

type User interface {
	Get() ([]dto.User, error)
}
type user struct {
	//ch chan dto.Page
}

func NewUser() User {
	return &user{}
}

func (u *user) Get() ([]dto.User, error) {
	ch := make(chan dto.Page)
	var wg sync.WaitGroup
	var page dto.Page
	resp, err := http.Get("https://reqres.in/api/users?page=1")
	if err != nil {
		log.Print("Not able to get the HTTP response ", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	error := json.Unmarshal(body, &page)
	if error != nil {
		log.Fatalln(error)
	}

	// log.Print( "Before go routines: ", page)
	time.Sleep(time.Second * 5)
	wg.Add(1)
	go readPages(page, ch)
	go func() {
		wg.Wait()
		log.Print("before closing the channel: ", page)
		close(ch)
	}()
	//go writePages(ch)
	log.Print(page.Users)
	return page.Users, error
}

func readPages(page dto.Page, ch chan<- dto.Page) {

	log.Print(" Inside read pages")
	ch <- page
	time.Sleep(time.Second * 2)
	log.Print("after read channels: ")
}

// func writePages(out <-chan dto.Page) (pages []dto.Page) {
// 	log.Print("writePages ", len(out))
// 	for i := 0; i < len(out); i++ {
// 		page := <-out
// 		log.Print(" channel read ", page)
// 		pages = append(pages, page)

// 	}
// 	return pages

// }
