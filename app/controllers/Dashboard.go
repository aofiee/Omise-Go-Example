package controllers

import (
	"github.com/revel/revel"
)

// Dashboard structure
type Dashboard struct {
	*revel.Controller
}

// Index method
func (c Dashboard) Index() revel.Result {
	return c.Render()
}
