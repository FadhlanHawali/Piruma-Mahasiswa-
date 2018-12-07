package model

type Mahasiswa struct {
	ID uint `gorm:"primary_key"`
	IdMahasiswa string `json:"id_mahasiswa"`
	Username string `json:"username"`
	Email string `json:"email"`
	NIM string `json:"nim"`
	Telepon int64 `json:"telepon"`
	Password string `json:"password"`
}


type History struct {
	ID uint `gorm:"primary_key"`
	IdMahasiswa string `json:"id_mahasiswa"`
	IdPemesanan string `json:"id_pemesanan"`
	Ruangan string `json:"ruangan"`
	Departemen string `json:"departemen"`
	PenanggungJawab string `json:"penanggung_jawab"`
	Telepon string `json:"telepon"`
	Keterangan string `json:"keterangan"`
	TimestampStart string `json:"timestamp_start"`
	TimestampEnd string `json:"timestap_end"`
}

type Bank struct {
	ID uint `gorm:"primary_key"`
	Name string `json:"name"`
	Saldo int `json:"saldo"`
}