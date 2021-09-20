package v1

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"picklejar/database"
	"picklejar/models"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func GetUser(c *gin.Context) {
	if database.Err != nil {
		panic(database.Err)
	}

	var user models.User
	id := c.Param("id")

	u64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		panic(err)
	}
	result := uint(u64)

	user.ID = result

	database.Db.Where("id = ?", user.ID).First(&user)
	c.JSON(http.StatusOK, user)

}

func RegisterUser(user models.User) {
	if database.Err != nil {
		panic(database.Err)
	}

	database.Db.Create(&user)
}

func CreateUser(c *gin.Context) {

	if database.Err != nil {
		panic(database.Err)
	}

	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		panic(err)
	}

	var user models.User

	json.Unmarshal(jsonData, &user)

	database.Db.Create(&user)

	c.JSON(http.StatusOK, user)

}

// soft deletes row
func DeleteUser(c *gin.Context) {
	if database.Err != nil {
		panic(database.Err)
	}
	var user models.User
	id := c.Param("id")

	u64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		panic(err)
	}
	result := uint(u64)

	user.ID = result

	database.Db.Where("id = ?", result).Delete(&user)

	c.JSON(http.StatusOK, user.ID)
}

func UpdateUser(c *gin.Context) {
	if database.Err != nil {
		panic(database.Err)
	}

	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		panic(err)
	}

	var user models.User

	id := c.Param("id")

	u64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		panic(err)
	}
	result := uint(u64)

	database.Db.Where("id = ?", result).Find(&user)
	json.Unmarshal(jsonData, &user)
	database.Db.Save(&user)

	c.JSON(http.StatusOK, user)
}
