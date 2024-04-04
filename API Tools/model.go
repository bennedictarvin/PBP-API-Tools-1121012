package main

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Berita struct {
	ID      int    `json:"id"`
	Tanggal string `json:"tanggal"`
	Judul   string `json:"judul"`
	Isi     string `json:"isi"`
}
