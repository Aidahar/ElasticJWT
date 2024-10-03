package domain

type Answer struct {
	Index     string  `json:"index"`
	Total     int     `json:"total"`
	Places    []Store `json:"places"`
	Prev_Page int     `json:"prev_page"`
	Next_Page int     `json:"next_page"`
	Last_page int     `json:"last_page"`
}

type Store struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Addresses string `json:"addresses"`
	Phone     string `json:"phone"`
	Location
}

type Location struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type Token struct {
	TokenStr string `json:"token"`
}
