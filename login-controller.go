package controllers

import (
	"github.com/ScottHuangZL/gin-jwt-session"
	"github.com/gin-gonic/gin"
  "picklejar/models"
	"picklejar/utils"
	"fmt"
	"net/http"
	"time"
)

//LoginHandler for login page , it also can use for logout since it delete all stored session
func LoginHandler(c *gin.Context) {
	flashes := session.GetFlashes(c)
	session.DeleteAllSession(c)

	c.HTML(http.StatusOK, "/", gin.H{
		"title":   "Dashboard",
		"flashes": flashes,
	})

  /*
  c.Redirect(http.StatusOK, "/Login", gin.H{
    "flashes": flashes,
  })*/
}

/*
//HomeHandler is the home handler
//will show home page, also according login/logout action to navigate
func HomeHandler(c *gin.Context) {
	// action := strings.ToLower(c.Param("action"))
	// path := strings.ToLower(c.Request.URL.Path)

	flashes := session.GetFlashes(c)
	username, err := session.ValidateJWTToken(c)
	loginFlag := false
	if err == nil && username != "" {
		loginFlag = true
	}
	c.HTML(http.StatusOK, "/", gin.H{
		"title":     "Main website",
		"now":       time.Now(),
		"flashes":   flashes,
		"loginFlag": loginFlag,
		"username":  username,
	})
}*/

//ValidateJwtLoginHandler validate the login and redirect to correct link
func ValidateLoginHandler(c *gin.Context) {
	var form models.LoginRequest
	//try get login info



	if err := c.ShouldBind(&form); err != nil {
		session.SetFlash(c, "Get login info error: "+err.Error())
		c.Redirect(http.StatusMovedPermanently, "/login")
		return
	}
	fmt.Printf("Email: %s, Password: %s", form.User.Email, form.User.Password)
	//validate login info
	if ok := utils.VerifyUserLogin(form.User.Email, form.User.Password); !ok {
		session.SetFlash(c, "Error : username or password")
		c.Redirect(http.StatusMovedPermanently, "/login")
		return
	}
	//login info is correct, can generate JWT token and store in clien side now
	tokenString, err := session.GenerateJWTToken(form.User.Email, time.Hour*time.Duration(1))
	if err != nil {
		session.SetFlash(c, "Error Generate token string: "+err.Error())
		c.Redirect(http.StatusMovedPermanently, "/login")
		return
	}

	err = session.SetTokenString(c, tokenString, 60*60) //60 minutes
	if err != nil {
		session.SetFlash(c, "Error set token string: "+err.Error())
		c.Redirect(http.StatusMovedPermanently, "/login")
		return
	}
	session.SetFlash(c, "success : successful login")
	session.SetFlash(c, "username : "+form.User.Email)
//	c.Redirect(http.StatusMovedPermanently, "http://192.168.1.63:3000/")
c.JSON(http.StatusOK, gin.H{"message": "User Sign In successfully",})
	return
}
