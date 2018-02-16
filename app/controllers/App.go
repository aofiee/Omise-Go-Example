package controllers

import (
	"fmt"
	"net/http"
	"regexp"
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
func (c App) Login(username string, password string, remember string) revel.Result {
	// delete(c.Session, "foo")
	// fmt.Println(c.Session["foo"])

	c.Validation.Required(username)
	c.Validation.MaxSize(username, 15)
	c.Validation.MinSize(username, 4)
	c.Validation.Match(username, regexp.MustCompile("^\\w*$"))

	c.Validation.Required(password)
	c.Validation.MinSize(password, 4)
	c.Validation.Match(password, regexp.MustCompile("^\\w*$"))

	if c.Validation.HasErrors() {
		c.Flash.Error("กรุณากรอก Username และ Password โดยมีความยาวตั้งแต่ 4 - 15 ตัวอักษร")
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Index)
	}

	db := models.Gorm
	println("Gorm works? : " + strconv.FormatBool(db != nil))
	var user models.User
	db.Where("username = ?", username).First(&user)
	if user.Username != "" && user.Password != "" {
		if models.CheckPasswordHash(password, user.Password) {
			fmt.Println(user.Username)
			if remember == "on" {
				newCookie := &http.Cookie{Name: "rememberLogin", Value: "on"}
				c.SetCookie(newCookie)
			}
		}
	}

	return c.Render(username, password)
}
