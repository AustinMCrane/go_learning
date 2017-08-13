package models

import (
	"os"

	"github.com/austinmcrane/wedding_app_api/dbutils"
)

var DB, dbError = dbutils.Init(os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

func init() {

}
