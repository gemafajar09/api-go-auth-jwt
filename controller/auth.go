package controller

import (
	token "api-go-auth-jwt/jwt"
	"api-go-auth-jwt/model"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type DataUser struct {
	Nama     string `json:"nama"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(c *gin.Context) {
	// deklarikan type
	var input DataUser
	// cek apakah respon dari json ada atau tidak
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// generate hash password
	password, _ := model.PasswordHash(input.Password)
	// query untuk data yang akan di inputkan ke database
	user := model.User{Nama: input.Nama, Username: input.Username, Password: password}
	// perintah buka koneksi database
	db := c.MustGet("db").(*gorm.DB)
	// ekseskusi perintah ke dalam bentuk sql
	db.Create(&user)
	// cek data dari inputan
	c.JSON(http.StatusOK, gin.H{"data": user})
}

type LoginUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Logins(c *gin.Context) {
	// variable error
	var err error
	// ambil dari type yg sudah di deklarasikan
	var user LoginUser
	// deklarasi data model user
	users := model.User{}
	// cek apakah error atau sukses
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// inisial data username dari json inputan
	username := user.Username
	password := user.Password

	// buka konseksi ke database
	db := c.MustGet("db").(*gorm.DB)
	// cek apakah username ada
	eror := db.First(&users, "username = ?", username).Error
	// jika username tidak ada maka return pesan
	if eror != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username Salah"})
		return
	}
	// cek apakah password betul atau salah
	err = model.PasswordCek(password, users.Password)
	// jika password salah return pesan error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password Salah"})
		return
	}
	// generate token
	token, err := token.GenerateToken(users.Id)
	// return token via json
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func UserId(c *gin.Context) {
	// panggil model user
	users := model.User{}
	// cek token yang dikirim dari header
	Id, err := token.ExtractTokenID(c)
	// jika token yg di kirim salah return pesan
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// buka koneksi ke database
	db := c.MustGet("db").(*gorm.DB)
	// cek jika token benar periksa data user
	if err := db.First(&users, "id = ?", Id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// return data json berdasarkan data user
	c.JSON(http.StatusOK, gin.H{"message": "Berhasil", "data": users})
}
