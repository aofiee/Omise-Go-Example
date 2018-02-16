package controllers

import (
	"fmt"
	"strings"
	"time"

	"github.com/aofiee666/OmiseWallet/app/models"

	"github.com/revel/revel"
)

// Dashboard structure
type Dashboard struct {
	*revel.Controller
	App
}

var (
	myName string
)

// Index method
func (c Dashboard) Index() revel.Result {
	myName := strings.Title(c.Session["username"])
	return c.Render(myName)
}

//checkUser func
func (c Dashboard) checkUser() revel.Result {
	if user := c.connected(); user == nil {
		c.Flash.Error("Please log in before")
		return c.Redirect(App.Index)
	}
	return nil
}

//Logout func
func (c Dashboard) Logout() revel.Result {
	for k := range c.Session {
		delete(c.Session, k)
	}
	c.ViewArgs["username"] = nil
	return c.Redirect(App.Index)
}

// PublicKey func
func (c Dashboard) PublicKey() revel.Result {
	myName := strings.Title(c.Session["username"])
	db := models.Gorm
	var omise models.OmiseKey
	db.First(&omise)
	publickey := omise.PublicKey
	secretkey := omise.SecretKey
	return c.Render(myName, publickey, secretkey)
}

// UpdateKey func
func (c Dashboard) UpdateKey(publickey string, secretkey string) revel.Result {
	fmt.Println(publickey, secretkey)
	c.Validation.Required(publickey)
	c.Validation.Required(secretkey)
	if c.Validation.HasErrors() {
		c.Flash.Error("กรุณากรอก public key และ secret key ด้วยครับ")
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(Dashboard.PublicKey)
	}
	myName := strings.Title(c.Session["username"])
	db := models.Gorm
	var omise models.OmiseKey
	db.First(&omise)

	if omise.ID == 0 {
		db.FirstOrCreate(&omise, models.OmiseKey{
			PublicKey:   publickey,
			SecretKey:   secretkey,
			CreatedDate: time.Now(),
		})
	} else {
		omise.PublicKey = publickey
		omise.SecretKey = secretkey
		db.Save(&omise)
	}
	c.ViewArgs["myName"] = myName
	c.ViewArgs["publickey"] = publickey
	c.ViewArgs["secretkey"] = secretkey
	return c.RenderTemplate("Dashboard/PublicKey.html")
}
