package controllers

import (
	"fmt"
	"strings"

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

	var d Dashboard
	p, s := d.getPublicAndSecretKey()
	fmt.Println(p, s)
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
