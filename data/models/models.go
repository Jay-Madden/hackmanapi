package models

type User struct {
	Id     int
	Name   string
	ApiKey string
}

type Requests struct {
	Id           int
	UserId       int
	ReturnedWord string
	Length       *string
}
