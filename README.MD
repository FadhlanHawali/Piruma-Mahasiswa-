# Piruma

Piruma adalah sebuah aplikasi peminjaman ruangan yang dibuat dalam rangka menyelesaikan tugas Interoperabilitas.

## Getting Started

Berikut adalah route dalam menggunakan API Piruma. Base URL yang digunakan adalah : https://piruma.au-syd.mybluemix.net

### Sign Up

Membuat akun baru

```
Route : /api/signup
Method : /POST
Body :
{
  "username":string,
  "email":string,
  "nim":string,
  "telepon":int64,
  "password":string",
  "confirm_password":string
}

Response :
{
  "status":"success"
}
```

### Login

Masuk ke akun yang sudah dibuat
```
Route : /api/login
Method : /POST
Body :
{
  "email":string,
  "password":string
}

Response :
{
  "status":"Success",
  "token":string
}
```
Token akan dipakai diseluruh operasi setelah login.

### List History

Untuk menampilkan list history dari pemesanan secara kasar (tidak detail).

```
Route : /api/list_history
Method : /GET
Header (
 Key = "Authorization"
 Value = token
)

Response :
{
    "count": int,
    "result": [
        {
            "id_pemesanan": string,
            "ruangan": string,
            "departemen": string,
            "status_peminjaman": string,
            "status_surat": string,
            "timestamp_start": string,
            "timestamp_end": string
        }
    ]
}
```
### Detail History

Untuk menampilkan detail dari history sebuah pemesanan ruangan

```
Route : /api/history?idPemesanan={id_pemesanan}
Method : /GET
Header (
 Key = "Authorization"
 Value = token
)

Response :
{
    "ID": string,
    "id_mahasiswa": string,
    "id_pemesanan": string,
    "ruangan": string,
    "departemen": string,
    "penanggung_jawab": string,
    "telepon": string,
    "keterangan": string,
    "timestamp_peminjaman": string,
    "status_peminjaman": string,
    "status_surat": string,
    "timestamp_start": string,
    "timestap_end": string
}
```


## Authors

* **Fadhlan Hawali** - *Back End Developer*
* **M. Reyzaldy Indra** - *Android Developer*
* **Faizal P.** - *Web Developer*
* **Raisul Rusdianawati** - *UI/UX Designer*