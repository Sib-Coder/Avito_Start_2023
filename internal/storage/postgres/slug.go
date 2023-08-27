package postgres

import (
	"Avito_Start_2023/internal/model"
	"fmt"
)

func (db *Database) DeleteSlug(name string) (bool, error) {
	_, err := db.db.Exec("DELETE FROM slug where name_slug = $1;", name)
	if err != nil {
		return false, err
	} else {
		return true, err
	}
}

func (db *Database) CreateSlug(name string) (bool, error) {
	_, err := db.db.Exec("INSERT INTO  slug ( name_slug) VALUES ($1);", name)
	if err != nil {
		return false, err
	}
	return true, err
}

// TODO наложить уникальность на имя slug и на сочетание id_slug и id_user
func (db *Database) CreateRelation(iduser string, idslug string) (bool, error) {
	_, err := db.db.Exec("INSERT INTO slugtraker (id_user, id_slug) VALUES ($1,$2);", iduser, idslug)
	if err != nil {
		return false, err
	}
	return true, err
}

func (db *Database) ExecSlugNamesUser(iduser string) ([]model.Slug, error) {
	var slug model.Slug
	var slugs []model.Slug
	res, err := db.db.Query("SELECT name_slug from slugtraker JOIN public.slug s on s.id_slug = slugtraker.id_slug where id_user =$1;", iduser)
	if err != nil {
		return nil, err
		fmt.Println(err)
	}
	for res.Next() {
		err = res.Scan(&slug.Name)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		slugs = append(slugs, slug)
	}
	return slugs, nil
}
