package controllers

import (
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
	return c.Render(username, password)
}
