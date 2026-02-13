package siswa

import (
	"time"

	"gorm.io/gorm"
)

// Siswa represents the student model.
type Siswa struct {
	ID           uint           `gorm:"primarykey" json:"id"`
	NIS          string         `gorm:"unique;not null" json:"nis"` // Nomor Induk Siswa
	Nama         string         `gorm:"not null" json:"nama"`
	JenisKelamin string         `gorm:"not null" json:"jenis_kelamin"` // L / P
	TempatLahir  string         `json:"tempat_lahir"`
	TanggalLahir *time.Time     `json:"tanggal_lahir"`
	Alamat       string         `json:"alamat"`
	NoTelepon    string         `json:"no_telepon"`
	Email        string         `gorm:"unique" json:"email"`
	Kelas        string         `json:"kelas"`
	TahunMasuk   int            `json:"tahun_masuk"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}
