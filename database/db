package database

import (
	"picklejar/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var Db *gorm.DB
var Err error

func ConnectDatabase() {
	Db, Err = gorm.Open("mysql", "admin:admin123!@tcp(picklejar-mysql.cereiuk6anjz.us-east-2.rds.amazonaws.com:3306)/picklejar?parseTime=true")

	if Err != nil {
		panic(Err)
	}

	Db.AutoMigrate(&models.User{})
	//Db.AutoMigrate(&models.Log{})

}

/*

func VerifyUserLogin(email, password string) bool {
  if database.Err != nil {
      panic(database.Err)
  }

  var user models.User

  database.Db.Where("email = ? AND password = ?", email, password).First(&user)

  fmt.Printf("email: %s", user.Email)

  if(user.Email != nil) {
    return true
  }

  return false
}
*/
