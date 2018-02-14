package controllers

import (
	"github.com/revel/revel"
)

type Applogin struct {
	*revel.Controller
}

func (c Applogin) Index() revel.Result {
	return c.Render()
}

func (c Applogin) Login(username string, password string) revel.Result {
	return c.Render(username, password)
}
