package main

import (
	"api-go-auth-jwt/api"
	"api-go-auth-jwt/model"

	"api-go-auth-jwt/router"
)

func main() {

	db := api.Koneksi()
	db.AutoMigrate(&model.User{})

	r := router.RouterAlamat(db)
	r.Run(":3030")
}
