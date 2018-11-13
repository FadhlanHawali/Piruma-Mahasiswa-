package model

type Login struct {
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type SignUp struct {
	Username string `json:"username" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	NIM string `json:"nim" binding:"required"`
	Telepon int64 `json:"telepon,string" binding:"required"`
	Password string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
}

type AddHistory struct {
	Ruangan string `json:"ruangan" binding:"required"`
	Departemen string `json:"departemen" binding:"required"`
	PenanggungJawab string `json:"penanggung_jawab" binding:"required"`
	Telepon string `json:"telepon" binding:"required"`
	Keterangan string `json:"keterangan" binding:"required"`
	TimestampPeminjaman string `json:"timestamp_peminjaman" binding:"required"`
	StatusPeminjaman string `json:"status_peminjaman" binding:"required"`
	StatusSurat string `json:"status_surat" binding:"required"`
	TimestampStart string `json:"timestamp_start" binding:"required"`
	TimestapEnd string `json:"timestamp_end" binding:"required"`
}

type ListHistory struct {
	IdPemesanan string `json:"id_pemesanan"`
	Ruangan string `json:"ruangan"`
	Departemen string `json:"departemen"`
	StatusPeminjaman string `json:"status_peminjaman"`
	StatusSurat string `json:"status_surat"`
	TimestampStart string `json:"timestamp_start"`
	TimestapEnd string `json:"timestamp_end"`

}
type SearchRuangan struct {
	Kapasitas string `json:"kapasitas" binding:"required"`
	TimeStamp TimeStamp
}

type TimeStamp struct {
	TimestampStart string `json:"timestamp_start" binding:"required"`
	TimestampEnd string `json:"timestamp_end" binding:"required"`
}



