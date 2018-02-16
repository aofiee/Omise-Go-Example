package models

import (
	"regexp"
	"time"

	"github.com/revel/revel"
)

// User Structure
type User struct {
	ID          int64
	Username    string
	Password    string
	Role        int
	CreatedDate time.Time
}

// Validate Users Data
func (u *User) Validate(v *revel.Validation) {
	v.Required(u.Username)
	v.Required(u.Password)
	v.Required(u.Role)
	v.MaxSize(u.Username, 15).MessageKey("Username must be between 4-15 characters long")
	v.MinSize(u.Username, 4).Message("Username must be between 4-15 characters long")
	v.MinSize(u.Password, 4).Message("Password must be beless than 4 characters long")
	v.Match(u.Username, regexp.MustCompile("^\\w*$"))
	v.Match(u.Password, regexp.MustCompile("^\\w*$"))
	return
}
