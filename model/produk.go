package model

type Produk struct {
	Id         uint   `json:"id"`
	NamaProduk string `json:"nama_produk"`
	Deskripsi  string `json:"deskripsi"`
	Harga      string `json:"harga"`
	Foto       string `json:"foto"`
}
