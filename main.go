package main

import (
	"jwt-client/auth"
	"jwt-client/controller"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kataras/golog"
	"github.com/subosito/gotenv"
)

func init() {
	// if len(os.Args) > 1 {
	// 	gotenv.Load(os.Args[1])
	// }
	if e := gotenv.Load(".env"); e != nil {
		log.Fatal(e)
	}
}

func main() {
	r := gin.Default()

	auth, e := auth.MiddlewareJWT()
	if e != nil {
		log.Fatal("error: " + e.Error())
	}

	r.POST("/login", auth.LoginHandler)
	r.POST("/logout", auth.LogoutHandler)

	user := r.Group("/user")
	user.Use(auth.MiddlewareFunc())
	{
		user.GET("/", controller.GetUser)
		user.GET("/:id", controller.GetByIDUser)
		user.POST("/create", controller.CreateUser)
		user.PUT("/:id", controller.UpdateUser)
		user.DELETE("/:id", controller.DeleteUser)
	}

	if e = r.Run(":" + os.Getenv("PORT")); e != nil {
		golog.Error(e)
	}

}
