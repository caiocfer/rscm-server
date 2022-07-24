package db

import (
	"database/sql"
	"rscm/src/config"
)

func Connect() (*sql.DB, error) {
	DBMS := "mysql"
	db, error := sql.Open(DBMS, config.DBConnection)
	if error != nil {
		return nil, error
	}

	if error = db.Ping(); error != nil {
		db.Close()
		return nil, error
	}

	return db, nil

}
