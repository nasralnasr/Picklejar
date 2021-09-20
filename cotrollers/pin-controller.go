package controllers


import (
  "picklejar/models"
  "picklejar/utils"
  "github.com/gin-gonic/gin"
  "net/http"
)

func ValidatePinHandler(c *gin.Context){
  var form models.PinRequest

  c.ShouldBind(&form)


  if(utils.IsPinValid(form.RaspPin, form.UserID)){
    c.JSON(http.StatusOK, "valid")
  } else {
    c.JSON(http.StatusBadRequest, "invalid")
  }
}
