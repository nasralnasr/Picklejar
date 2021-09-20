package utils

import (
	"log"
	"picklejar/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

/*
func ValidateUser(email, password string) bool {

	if email == "admin@picklejar.com" && password == "123" ||
		email == "user1" && password == "user1" ||
		email == "user2" && password == "user2" {
		return true
	}
	return false


	return true
}
*/

var db *gorm.DB
var err error

func connectDatabase() {
	db, err = gorm.Open("mysql", "admin:admin123!@tcp(picklejar-mysql.cereiuk6anjz.us-east-2.rds.amazonaws.com:3306)/picklejar?parseTime=true")

	if err != nil {
		panic(err)
	}

}

func DoesPhoneExist(phoneNumber string) bool {
	connectDatabase()

	if err != nil {
		panic(err)
	}

	var user models.User

	db.Where("phone_number = ?", phoneNumber).First(&user)

	log.Printf("phoneNumber: %s", user.PhoneNumber)

	if len(user.Email) > 0 {
		return true
	}

	return false
}

func IsPinValid(pin, id string) bool {

	connectDatabase()

	if err != nil {
		panic(err)
	}

	var user models.User
	//var pinToCheck string = pin
//  var idToCheck string = id

	db.Where("pin = ? AND id = ?", pin, id).First(&user)

	log.Printf("pin: %s", pin)

	if len(user.Email) > 0 {
		return true
	}

	return false

}

func VerifyUserLogin(email, password string) bool {

	connectDatabase()

	log.Printf("email: %s, password: %s", email, password)

	if err != nil {
		panic(err)
	}

	var user models.User

	db.Where("email = ? AND password = ?", email, password).First(&user)

	log.Printf("email: %s", user.Email)

	if len(user.Email) > 0 {
		return true
	}

	return false
}

/*
func NewUserLog(phoneNumber string) {
	connectDatabase()

	var logs models.Log
	logs.PhoneNumber = phoneNumber
	log.Printf("Creating new log with phonenumber: %s", phoneNumber)

	if err != nil {
		panic(err)
	}
	db.Create(&logs)
}*/
