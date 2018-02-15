package controllers

import (
	"time"

	"github.com/aofiee666/OmiseWallet/app/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/revel/revel"
)

// App structure
type Gorm struct {
	*revel.Controller
	db *gorm.DB
}

var Gdb *gorm.DB

func InitDB() {

	Gdb, err := gorm.Open("mysql", "root:root@/wallet?charset=utf8&parseTime=True")
	if err != nil {
		panic("Unable to connect to the database")
	}
	Gdb.DB().Ping()
	Gdb.DB().SetMaxIdleConns(10)
	Gdb.DB().SetMaxOpenConns(100)
	Gdb.SingularTable(true)

	if !Gdb.HasTable(&models.User{}) {
		Gdb.CreateTable(&models.User{})
		user := models.User{
			Username:    "admin",
			Password:    "password",
			CreatedDate: time.Now(),
		}
		Gdb.Create(&user)
	}
}
