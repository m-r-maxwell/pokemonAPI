package cfg

import (
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func InitConfig() *sqlx.DB {
	host := "localhost"
	port := "5432"
	var user string
	var password string
	// if we cant get the env variable, we will use the default value
	if os.Getenv("db_user") == "" {
		user = "postgres"
	} else {
		user = os.Getenv("db_user")
	}
	if os.Getenv("db_password") == "" {
		password = "password"
	} else {
		password = os.Getenv("db_password")
	}
	dbname := "pokemon"
	sslmode := "disable"
	db, err := sqlx.Connect("postgres", "host="+host+" port="+port+" user="+user+" password="+password+" dbname="+dbname+" sslmode="+sslmode)
	if err != nil {
		panic(err)
	}
	return db

}
