package dbutils

import (
	"database/sql"
	"fmt"
)

func Init(host string, port string, username string, password string, dbname string) (*sql.DB, error) {
	c := fmt.Sprintf("user=%s dbname=%s port=%s password=%s host=%s sslmode=disable", username, dbname, port, password, host)
	db, err := sql.Open("postgres", c)
	if err != nil {
		return nil, err
	}
	return db, err
}
