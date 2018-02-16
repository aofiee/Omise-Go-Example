package controllers

import "github.com/revel/revel"

func init() {
	revel.InterceptMethod(Dashboard.checkUser, revel.BEFORE)
}
