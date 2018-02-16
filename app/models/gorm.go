package models

import (
	"net/url"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //
	"github.com/revel/revel"
	"golang.org/x/crypto/bcrypt"
)

// Gorm var
var (
	Gorm *gorm.DB
)

// InitDB var
func InitDB() {
	driver := revel.Config.StringDefault("db.driver", "mysql")
	username := revel.Config.StringDefault("db.username", "root")
	password := revel.Config.StringDefault("db.password", "root")
	protocol := revel.Config.StringDefault("db.protocol", "tcp")
	hostname := revel.Config.StringDefault("db.hostname", "localhost")
	port := revel.Config.StringDefault("db.port", "3306")
	dbName := revel.Config.StringDefault("db.database_name", "wallet")
	timezone := revel.Config.StringDefault("db.timezone", "UTC")
	charset := revel.Config.StringDefault("db.charset", "utf8")
	collation := revel.Config.StringDefault("db.collation", "utf8_bin")

	dsn := username + ":" + password +
		"@" + protocol + "(" + hostname + ":" + port + ")" +
		"/" + dbName +
		"?parseTime=True&loc=" + url.PathEscape(timezone) +
		"&charset=" + charset + "&collation=" + collation

	var err error
	Gorm, err = gorm.Open(driver, dsn)
	if err != nil {
		defer Gorm.Close() // panic
		panic(err)
	}

	Gorm.LogMode(true)
	Gorm.SetLogger(gorm.Logger{revel.INFO})
	Gorm.SingularTable(true)

	autoMigrate()
	revel.INFO.Println("Connected to database.")
}

func autoMigrate() {
	Gorm.Set("gorm:table_options", "ENGINE=InnoDB")
	if !Gorm.HasTable(&User{}) {
		Gorm.AutoMigrate(&User{})
		HashPassword, _ := HashPassword("password")
		user := User{
			Username:    "admin",
			Password:    HashPassword,
			Role:        1, //1 = admin 0 = default
			CreatedDate: time.Now(),
		}
		Gorm.Create(&user)
	}
	if !Gorm.HasTable(&OmiseKey{}) {
		Gorm.AutoMigrate(&OmiseKey{})
	}
}

// HashPassword func
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash func
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
