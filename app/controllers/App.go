package controllers

import (
	"fmt"
	"strconv"

	"github.com/aofiee666/OmiseWallet/app/models"
	"github.com/revel/revel"
)

// App structure
type App struct {
	*revel.Controller
}

// Index method
func (c App) Index() revel.Result {
	return c.Render()
}

// Login metthod
func (c App) Login(username string, password string) revel.Result {
	db := models.Gorm
	println("Gorm works? : " + strconv.FormatBool(db != nil))
	var user []models.User
	db.Where("username = ?", "admin").Find(&user)
	fmt.Println(user)
	return c.Render(username, password)
}
