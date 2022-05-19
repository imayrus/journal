package models

type Entry []struct {
	Entry string `json:"entry"`
}

type LoginDetails struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	ID       int    `json:"id"`
}
