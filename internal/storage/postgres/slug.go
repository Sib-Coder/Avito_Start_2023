package postgres

import (
	"avitoStart/internal/model"
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

// TODO Если не надо удалить эту функцию
func (db *Database) ExecIdSlug(name string) (string, error) {
	var id string
	res, err := db.db.Query("SELECT id_slug from slug where name_slug =$1;", name)
	if err != nil {
		return " ", err
	}
	for res.Next() {
		err = res.Scan(&id)
		if err != nil {
			return " ", err
		}
		//проверка на существование
	}
	return id, err
}

func (db *Database) DeleteRelation(iduser string, name string) (bool, error) {
	_, err := db.db.Exec("DELETE FROM slugtraker WHERE id_user = $1 AND id_slug IN (SELECT slug.id_slug from slug where name_slug=$2);", iduser, name)
	if err != nil {
		return false, err
	}
	return true, err
}

func (db *Database) CreateRelation(iduser string, name string) (bool, error) {
	tx, err := db.db.Beginx()
	if err != nil {
		fmt.Println(err)
	}

	stmt, err := tx.Prepare("SELECT id_slug from slug where name_slug =$1;")
	if err != nil {
		fmt.Println("error: %v\n", err)
	}
	id_slug := ""
	err = stmt.QueryRow(name).Scan(&id_slug)
	if err != nil {
		tx.Rollback()
		fmt.Println(err)
	}
	//fmt.Println("ID_SLUG:", id_slug)
	_, err = tx.Exec("INSERT INTO slugtraker (id_user, id_slug) VALUES ($1,$2);", iduser, id_slug)
	if err != nil {
		tx.Rollback()
		fmt.Println(err)
	}

	err = tx.Commit()
	if err != nil {
		fmt.Println(err)
		return false, err
	} else {
		return true, err
	}
}
