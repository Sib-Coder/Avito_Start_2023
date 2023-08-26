package main

import (
	"Avito_Start_2023/internal/storage/postgres"
	"fmt"
)

func main() {
	dsn := "host='127.0.0.1' port=5432 user=postgres password=postgres dbname=postgres sslmode=disable"
	//dsn:= TODO перенести в env
	dan, err := postgres.New(dsn)
	fmt.Println(dan, err)
	fmt.Println(dan.ExtractUsers())

	//user := model.User{
	//	Name:     "Pipi",
	//	Sameinfo: "Cuscus",
	//}
	//res, err := dan.AddUser(user)
	//fmt.Println(res, err)
	//fmt.Println(dan.ExtractUsers())
	//fmt.Println(dan.DeleteUser("f5556271-5c81-465c-8343-6dbc864104f3"))
	fmt.Println(dan.CreateSlug("BlaxaMuxa"))
}
