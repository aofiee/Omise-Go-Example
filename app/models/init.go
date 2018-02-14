package models

import (
	"database/sql"
	"log"

	"github.com/coopernurse/gorp"
	_ "github.com/go-sql-driver/mysql" //
)

// Global database references
var db *sql.DB
var dbmap *gorp.DbMap

// Database settings
var dbName = "wallet"
var dbUser = "root"
var dbPwd = "root"

// InitDB Create database connection
func InitDB() {
	var err error

	db, err = sql.Open("mysql", dbUser+":"+dbPwd+"@tcp(127.0.0.1:3306)/"+dbName)
	dbmap = &gorp.DbMap{
		Db: db,
		Dialect: gorp.MySQLDialect{
			"InnoDB",
			"UTF8",
		},
	}

	if err != nil {
		log.Println("Failed to connect to database: ")
		log.Panic(err)
	} else {
		err = db.Ping()

		if err != nil {
			log.Println("Failed to ping database: ")
			log.Panic(err)
		} else {
			log.Println("Database connected.")
		}
	}

	_ = dbmap.AddTableWithName(OmiseKey{}, "omisekey").SetKeys(false, "ID")
	dbmap.CreateTablesIfNotExists()
}
