package router

import (
	"api-go-auth-jwt/controller"
	"api-go-auth-jwt/middleware"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func RouterAlamat(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	belumlogin := r.Group("/api")
	belumlogin.POST("/user", controller.Register)
	belumlogin.POST("/login", controller.Logins)

	sudahlogin := r.Group("/admin")
	sudahlogin.Use(middleware.Middleware())
	sudahlogin.GET("/user", controller.UserId)

	return r
}
