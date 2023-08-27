package main

import (
	"Avito_Start_2023/internal/model"
	"Avito_Start_2023/internal/storage/postgres"
	"fmt"
)

func main() {
	dsn := "host='127.0.0.1' port=5432 user=postgres password=postgres dbname=postgres sslmode=disable"
	//dsn:= TODO перенести в env
	dan, err := postgres.New(dsn)
	fmt.Println(dan, err)
	fmt.Println(dan.ExtractUsers())

	user := model.User{
		Name:     "Pipi",
		Sameinfo: "Cuscus",
	}
	res, err := dan.AddUser(user)
	fmt.Println(res, err)
	fmt.Println(dan.ExtractUsers())
	//fmt.Println(dan.CreateSlug("BlaxaMuxa"))
	fmt.Println(dan.ExecIdSlug("BlaxaMuxa"))
	//////////////////////////////////////////////////////

	id_user := "f769ad21-f3d6-4523-8a56-cd84172093bc"
	//id_slug := "1105bed0-3c51-435a-8b53-32f9b0f28219"
	//fmt.Println(dan.CreateRelation(id_user, id_slug))
	fmt.Println(dan.DeleteRelation(id_user, "BlaxaMuxa"))
	//fmt.Println(dan.ExecSlugNamesUser(id_user))
	//fmt.Println(dan.DeleteUser("f5556271-5c81-465c-8343-6dbc864104f3"))

}
