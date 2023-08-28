package model

type User struct {
	Id       string
	Name     string `json:"Name"`
	Sameinfo string `json:"Sameinfo"`
}
type Slug struct {
	Name string
}
type UserQuery struct {
	ID string `query:"id"`
}
