package postgres

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
