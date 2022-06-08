package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

var DB *gorm.DB

type User struct {
	Id       uint   `json:"id" gorm:"primary_key"`
	Nama     string `json:"nama"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func PasswordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func PasswordCek(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
