package routers

import (
	"auth_in_go/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRouters(r *gin.Engine) {
	r.POST("login", controllers.Login)
	r.POST("signup", controllers.SignUp)
	r.GET("home", controllers.Home)
	r.GET("premium", controllers.Premium)
	r.GET("logout", controllers.Logout)
}
