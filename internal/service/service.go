package service

import "Avito_Start_2023/internal/model"

type Database interface {
	//users
	AddUser(user model.User) (bool, error)
	DeleteUser(id string) (bool, error)
	ExtractUsers() ([]model.User, error)

	//slug
	DeleteSlug(name string) (bool, error)
	CreateSlug(name string) (bool, error)
	ExecSlugNamesUser(iduser string) ([]model.Slug, error)
	DeleteRelation(iduser string, name string) (bool, error)
	CreateRelation(iduser string, name string) (bool, error)
}
type Service struct {
	db Database
}

func New(db Database) *Service {
	return &Service{
		db: db,
	}
}
