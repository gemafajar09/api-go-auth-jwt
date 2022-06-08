package controller

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"api-go-auth-jwt/model"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Produkisi struct {
	NamaProduk string `json:"namaproduk"`
	Deskripsi  string `json:"deskripsi"`
	Harga      string `json:"harga"`
	Foto       string `json:"foto"`
}

func AddProduk(c *gin.Context) {
	// cek file foto dari inputan
	file, header, err := c.Request.FormFile("foto")
	// cek jika file ada atau tidak
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
		return
	}
	// deklarasi nama file
	filename := header.Filename
	// lokasi tempat file disimpan
	out, err := os.Create("public/" + filename)
	// cek apaka berhasil atau gagal
	if err != nil {
		log.Fatal(err)
	}
	//
	defer out.Close()
	// pindahkan file foto kedalam file tmp
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
	// deklarasikan data yang akan disimpan ke database
	data := model.Produk{NamaProduk: c.PostForm("nama_produk"), Deskripsi: c.PostForm("deskripsi"), Harga: c.PostForm("harga"), Foto: filename}
	// buka koneksi database
	db := c.MustGet("db").(*gorm.DB)
	// ekseskusi perintah ke dalam bentuk sql
	db.Create(&data)
	// return hasil yang disimpan ke database tadi
	c.JSON(http.StatusOK, gin.H{"filepath": data})
}
