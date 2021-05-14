package controller

import (
	"github.com/gin-gonic/gin"
)

// Setup usersのルーティング
func Setup(r *gin.RouterGroup) {
	users := r.Group("/users")
	{
		u := UserController{}
		users.GET("", u.Index)
		// users.POST("/login", u.Login)
		// users.GET("/:id", u.GetUser)
		users.POST("", u.CreateUser)
		// users.PATCH("/:id", u.UpdateUser)
		// users.DELETE("/:id", u.DeleteUser)
	}

	login := r.Group("/login")
	{
		l := LoginController{}
		login.POST("", l.LoginPost)
	}
}
