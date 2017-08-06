package models

import (
	"github.com/austinmcrane/wedding_api/dbutils"
)

var DB, dbError = dbutils.Init("localhost", "5432", "postgres", "password", "wedding_app")

func init() {

}
