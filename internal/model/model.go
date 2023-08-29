package model

type User struct {
	Id       string `json:"Id"`
	Name     string `json:"Name"`
	Sameinfo string `json:"Sameinfo"`
}
type Slug struct {
	Name string
}

//type UserQuery struct {
//	ID string `query:"id"`
//}

type MasterData struct {
	Id     string   `json:"Id"`
	MasAdd []string `json:"MasAdd"`
	MasDel []string `json:"MasDel"`
}
type SlugQuery struct {
	Name string `json:"Name"`
}
