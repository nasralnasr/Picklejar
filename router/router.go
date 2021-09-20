package router

import (
	"github.com/gin-gonic/gin"
  "picklejar/router/api/v1"
	"github.com/ScottHuangZL/gin-jwt-session"
	"picklejar/controllers"
)

func InitRouter() *gin.Engine {
  router := gin.Default()
	session.NewStore()
	session.DefaultFlashSessionName = "flashSession"
	session.DefaultSessionName = "YouCanChangeTheFlashName"
	router.Use(session.ClearMiddleware())

/*
	corsConfig := cors.DefaultConfig()


	corsConfig.AllowOrigins = []string{"http://13.59.118.188:3000"}
	// To be able to send tokens to the server.
	corsConfig.AllowCredentials = true

	// OPTIONS method for ReactJS
	corsConfig.AddAllowMethods("OPTIONS")

	// Register the middleware
	router.Use(cors.New(corsConfig))
*/
	router.Use(CORSMiddleware())

	router.Delims("{%", "%}")
	router.MaxMultipartMemory = 8 << 20 // 8 MiB

	router.GET("/", controllers.HomeHandler)

//	router.GET("/login", controllers.LoginHandler)
	router.GET("/logout", controllers.LoginHandler)
	router.POST("/validate-login", controllers.ValidateLoginHandler)
	router.POST("/validate-pin", controllers.ValidatePinHandler)

	router.POST("/register", controllers.RegisterHandler)
	router.OPTIONS("/register", controllers.RegisterHandler)
  groupV1 := router.Group("/api/v1")

  groupV1.GET("/user/:id", v1.GetUser)
  groupV1.POST("/user", v1.CreateUser)
  groupV1.PUT("/user/:id", v1.UpdateUser)
  groupV1.DELETE("/user/:id", v1.DeleteUser)


  return router
}

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")



        c.Next()
    }
}
