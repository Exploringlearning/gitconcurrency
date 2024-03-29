package services

import (
	"encoding/json"
	"fmt"
	"ginconcurrency/internal/app/dto"
	"io"
	"log"
	"net/http"
	"sync"
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
	ch := make(chan dto.Page, 2)
	var wg sync.WaitGroup
	var users []dto.User

	firstPage := getPage(1)
	go getUsersList(ch, &wg, &users)

	for i := 1; i <= firstPage.TotalPages; i++ {
		wg.Add(1)
		pageNo := i
		go sendPageToChannel(ch, &wg, pageNo)
	}
	wg.Wait()
	close(ch)

	log.Println("Final list:", users)
	return users, nil
}

func sendPageToChannel(ch chan<- dto.Page, wg *sync.WaitGroup, pageNo int) {
	defer func() {
		log.Println("Fetch page completed!")
		wg.Done()
	}()

	page := getPage(pageNo)
	log.Println("Page -> ", page)
	ch <- page

}

func getPage(pageNo int) dto.Page {
	resp, err := http.Get(fmt.Sprintf("https://reqres.in/api/users?page=%d", pageNo))
	if err != nil {
		log.Print("Not able to get the HTTP response ", err)
	}
	var page dto.Page
	body := getResponseBody(resp, err)
	serializeJsonToPageDto(body, &page)
	return page
}

func getUsersList(ch chan dto.Page, wg *sync.WaitGroup, users *[]dto.User) {

	for page := range ch {
		*users = append(*users, page.Users...)
	}

}

func serializeJsonToPageDto(body []byte, page *dto.Page) {
	if err := json.Unmarshal(body, &page); err != nil {
		log.Fatalln(err)
	}
}

func getResponseBody(resp *http.Response, err error) []byte {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	return body
}

//func readPages(page dto.Page, ch chan<- dto.Page) {
//	log.Print(" Inside read pages")
//	ch <- page
//	time.Sleep(time.Second * 2)
//	log.Print("after read channels: ")
//}

// func writePages(out <-chan dto.Page) (pages []dto.Page) {
// 	log.Print("writePages ", len(out))
// 	for i := 0; i < len(out); i++ {
// 		page := <-out
// 		log.Print(" channel read ", page)
// 		pages = append(pages, page)

// 	}
// 	return pages

// }
