package dto

type Page struct {
	Page      int `json:"page"`
	PerPage   int `json:"per_page"`
	Total     int `json:"total"`
	Totalpage int `json:"totalpage"`
	Users      []User `json:"data"`
}

type User struct {
	Id        int `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Avatar    string `json:"avatar"`
}
