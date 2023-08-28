package service

import "avitoStart/internal/model"

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

// work with users
func (s *Service) AddUser(user model.User) (bool, error) {
	res, err := s.AddUser(user)
	if err != nil {
		return res, err
	}
	return res, err
}

func (s *Service) DeleteUser(id string) (bool, error) {
	res, err := s.DeleteUser(id)
	if err != nil {
		return res, err
	}
	return res, err
}

func (s *Service) ExtractUsers() ([]model.User, error) {
	res, err := s.ExtractUsers()
	if err != nil {
		return res, err
	}
	return res, err
}

//work with slug

func (s *Service) DeleteSlug(name string) (bool, error) {
	res, err := s.DeleteSlug(name)
	if err != nil {
		return res, err
	}
	return res, err
}

func (s *Service) CreateSlug(name string) (bool, error) {
	res, err := s.CreateSlug(name)
	if err != nil {
		return res, err
	}
	return res, err
}

func (s *Service) ExecSlugNamesUser(iduser string) ([]model.Slug, error) {
	res, err := s.ExecSlugNamesUser(iduser)
	if err != nil {
		return res, err
	}
	return res, err
}

//big logical func
//DeleteRelation(iduser string, name string) (bool, error)
//CreateRelation(iduser string, name string) (bool, error)
