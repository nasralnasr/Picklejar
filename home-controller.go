package controllers

import (
  "github.com/gin-gonic/gin"
  "net/http"
)

//
func HomeHandler(c *gin.Context){
  c.Redirect(http.StatusTemporaryRedirect, "http://192.168.1.63:3000/")
  return
}
