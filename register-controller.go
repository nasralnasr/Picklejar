package controllers

import (
	"net/http"
	"picklejar/models"
	v1 "picklejar/router/api/v1"
	"picklejar/utils"

	"github.com/gin-gonic/gin"
)

func RegisterHandler(c *gin.Context) {

	var form models.RegisterRequest
	c.ShouldBind(&form)

	if utils.DoesPhoneExist(form.User.PhoneNumber) {
		c.JSON(http.StatusBadRequest, "Phone Number Already Exists")
		return
	}

	form.User.Pin = "1234"

	v1.RegisterUser(form.User)
	//utils.NewUserLog(form.User.PhoneNumber)

	utils.SubscribeToSMS(&form.User.PhoneNumber)
	utils.PublishSMS(&form.User.PhoneNumber, form.User.FirstName)
	c.JSON(http.StatusOK, form.User)
}
