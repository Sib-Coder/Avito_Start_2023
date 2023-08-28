package endpoint

import (
	"avitoStart/internal/model"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Service interface {
	//users
	AddUser(user model.User) (bool, error)
	DeleteUser(id string) (bool, error)
	ExtractUsers() ([]model.User, error)

	//slug
	DeleteSlug(name string) (bool, error)
	CreateSlug(name string) (bool, error)
	ExecSlugNamesUser(iduser string) ([]model.Slug, error)
}

type Endpoint struct {
	s Service
}

func New(s Service) *Endpoint {
	return &Endpoint{
		s: s,
	}
}

func (e *Endpoint) AddUser(ctx echo.Context) error {
	user := model.User{}

	err := ctx.Bind(&user)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	fmt.Println(user)
	res, err := e.s.AddUser(user)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, res)
}

func (e *Endpoint) DeleteUser(ctx echo.Context) error {
	var user model.UserQuery
	err := ctx.Bind(&user)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	fmt.Println(user)
	res, err := e.s.DeleteUser(user.ID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, res)
}
func (e *Endpoint) ExtractUsers(ctx echo.Context) error {
	res, err := e.s.ExtractUsers()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, res)
}

// пример идеи как буду работать с slug коррект
//type MyData struct {
//	MyArray []int `json:"myarray"`
//}
//
//func NotUse() {
//
//	jsonData := []byte(`{"myarray":[1, 2, 3, 4, 5]}`)
//	var data MyData
//	err := json.Unmarshal(jsonData, &data)
//	if err != nil {
//		fmt.Println("Error unmarshalling JSON:", err)
//		return
//	}
//	fmt.Println(data.MyArray)
//
//	my_mas := data.MyArray
//	fmt.Println(my_mas[1])
//
//	for k, el := range my_mas {
//		fmt.Println("key:", k, "element:", el)
//	}
//}
