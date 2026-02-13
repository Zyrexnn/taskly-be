package siswa

import "time"

// CreateSiswaRequestDTO defines the structure for creating a new siswa.
type CreateSiswaRequestDTO struct {
	NIS          string `json:"nis" validate:"required"`
	Nama         string `json:"nama" validate:"required"`
	JenisKelamin string `json:"jenis_kelamin" validate:"required,oneof=L P"`
	TempatLahir  string `json:"tempat_lahir"`
	TanggalLahir string `json:"tanggal_lahir"` // Format: YYYY-MM-DD
	Alamat       string `json:"alamat"`
	NoTelepon    string `json:"no_telepon"`
	Email        string `json:"email" validate:"omitempty,email"`
	Kelas        string `json:"kelas"`
	TahunMasuk   int    `json:"tahun_masuk"`
}

// UpdateSiswaRequestDTO defines the structure for updating a siswa.
type UpdateSiswaRequestDTO struct {
	NIS          string `json:"nis"`
	Nama         string `json:"nama"`
	JenisKelamin string `json:"jenis_kelamin" validate:"omitempty,oneof=L P"`
	TempatLahir  string `json:"tempat_lahir"`
	TanggalLahir string `json:"tanggal_lahir"` // Format: YYYY-MM-DD
	Alamat       string `json:"alamat"`
	NoTelepon    string `json:"no_telepon"`
	Email        string `json:"email" validate:"omitempty,email"`
	Kelas        string `json:"kelas"`
	TahunMasuk   int    `json:"tahun_masuk"`
}

// SiswaResponseDTO defines the structure for siswa data in responses.
type SiswaResponseDTO struct {
	ID           uint       `json:"id"`
	NIS          string     `json:"nis"`
	Nama         string     `json:"nama"`
	JenisKelamin string     `json:"jenis_kelamin"`
	TempatLahir  string     `json:"tempat_lahir"`
	TanggalLahir *time.Time `json:"tanggal_lahir"`
	Alamat       string     `json:"alamat"`
	NoTelepon    string     `json:"no_telepon"`
	Email        string     `json:"email"`
	Kelas        string     `json:"kelas"`
	TahunMasuk   int        `json:"tahun_masuk"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}
