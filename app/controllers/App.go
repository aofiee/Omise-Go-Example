package controllers

/* Package controllers for startup program
+-------------------------------------------------------------------------------------------+
|                                                                                           |
|                                                                                           |
|      								                                                        |
|                                                                                           |
|                                        .',,;,,..                                          |
|                                   ;d0NMMMMMMMMMMWXOo;.                                    |
|                                ;OWMMMMMMMMMMMMMMMMMMMMKo.                                 |
|                              cXMMMMMMMMMMMMMMMMMMMMMMMMMM0,                               |
|                            cXMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM0.                             |
|                          .KMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMWc                            |
|                          dMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM;                           |
|                          XMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMO                           |                           -
|                          XMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMX                           |
|                          XMMc:xXMMMMMMMMMMMMMMMMMMMMMMMNOo:cMMO                           |
|                          0MM,   .;dKWMMMMMMMMMMMMMMWKo'    'MM;                           |
|                          lMMK.      .;dXMMMMMMMMNkc.       0MK                            |
|                           OMMXl,.       'OMMMM0,       .'cKMMc                            |
|                            cNMMMMMNK0O0NMMXxNXMWX0O0KNWMMMMMX                             |
|                            :KMMMMMMMMMMMMK.kk.OMMMMMMMMMMMMMx                             |
|                            NMMMMMMXkdXMMMxoMMooMMMMMNOddXMMMW.                            |
|                            OMMMMk'   kMMMMMMMMMMMMMN    lMMO'                             |
|                             ;:..c.   OMMMMMWWMNoNMMx    .;.                               |
|                                      'MMkWMOoMKlMMW.                                      |
|                                       WMdXMklMNxMMx                                       |
|                                       WMd0MkcMWkMMo                                       |
|                                       NMd0MkcMNxMMl                                       |
|                                       XMd0Mk:MXdMM;                                       |
|                                       0Md0Mk:MOoMM.                                       |
|                                       lWoOMx;MloM0                                        |
|                                          ':. :..'                                         |
|                                                                                           |
+-------------------------------------------------------------------------------------------+
*/
import (
	"net/http"
	"regexp"
	"strconv"

	"Omise-Go-Example/app/models"

	"github.com/revel/revel"
)

// App structure
type App struct {
	*revel.Controller
}

// Index method
func (c App) Index() revel.Result {
	if c.connected() != nil {
		return c.Redirect(Dashboard.Index)
	}
	c.Flash.Error("Please log in first")
	return c.Render()
}

//connected func
func (c App) connected() *models.User {
	if c.ViewArgs["username"] != nil {
		return c.ViewArgs["username"].(*models.User)
	}
	if username, ok := c.Session["username"]; ok {
		return c.getUser(username)
	}
	return nil
}

//getUser func
func (c App) getUser(username string) (user *models.User) {
	user = &models.User{}
	db := models.Gorm
	db.Where("username = ?", username).First(&user)
	if user.Username != "" {
		return
	}
	return nil
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
			//fmt.Println(user.Role)
			if remember == "on" {
				newCookie := &http.Cookie{Name: "remember", Value: "on"}
				usernameCookie := &http.Cookie{Name: "username", Value: user.Username}
				passwordCookie := &http.Cookie{Name: "password", Value: password}
				c.SetCookie(newCookie)
				c.SetCookie(usernameCookie)
				c.SetCookie(passwordCookie)
				c.Session.SetNoExpiration()
			} else {
				newCookie := &http.Cookie{Name: "remember", Value: ""}
				usernameCookie := &http.Cookie{Name: "username", Value: ""}
				passwordCookie := &http.Cookie{Name: "password", Value: ""}
				c.SetCookie(newCookie)
				c.SetCookie(usernameCookie)
				c.SetCookie(passwordCookie)
				c.Session.SetDefaultExpiration()
			}
			c.Session["username"] = user.Username
			c.Session["role"] = string(user.Role)
			return c.Redirect(Dashboard.Index)

		} else {
			c.Flash.Error("Username และ Password ไม่ถูกต้อง")
			c.Validation.Keep()
			c.FlashParams()
			return c.Redirect(App.Index)
		}
	}

	return c.Render()
}

//getPublicAndSecretKey func
func (c App) getPublicAndSecretKey() (publickey string, secretkey string) {
	db := models.Gorm
	var omiseKey models.OmiseKey
	db.First(&omiseKey)
	publickey = omiseKey.PublicKey
	secretkey = omiseKey.SecretKey
	return publickey, secretkey
}
