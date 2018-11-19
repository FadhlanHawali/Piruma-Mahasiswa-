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
	StatusSurat string `json:"status_surat"`
	StatusPeminjaman bool `json:"status_peminjaman"`
	TimestampPeminjaman string `json:"timestamp_peminjaman"`
	TimestampStart string `json:"timestamp_start"`
	TimestapEnd string `json:"timestap_end"`
}

type Bank struct {
	ID uint `gorm:"primary_key"`
	Name string `json:"name"`
	Saldo int `json:"saldo"`
}