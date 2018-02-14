package controllers

import (
	"github.com/aofiee666/OmiseWallet/app/models"
	"github.com/revel/revel"
)

// Applogin structure
type Applogin struct {
	*revel.Controller
}

func init() {
	revel.OnAppStart(models.InitDB)
}

// Index method
func (c Applogin) Index() revel.Result {
	return c.Render()
}

// Login metthod
func (c Applogin) Login(username string, password string) revel.Result {
	return c.Render(username, password)
}
