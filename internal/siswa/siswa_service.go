package siswa

import (
	"errors"
	"tasklybe/internal/dto"
	"time"

	"gorm.io/gorm"
)

type Service interface {
	Create(req CreateSiswaRequestDTO) (*SiswaResponseDTO, error)
	GetAll(page, limit int, search string) (*dto.PaginatedResponse[SiswaResponseDTO], error)
	GetByID(id uint) (*SiswaResponseDTO, error)
	Update(id uint, req UpdateSiswaRequestDTO) (*SiswaResponseDTO, error)
	Delete(id uint) error
}

type service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) Service {
	return &service{db: db}
}

// Create creates a new siswa.
func (s *service) Create(req CreateSiswaRequestDTO) (*SiswaResponseDTO, error) {
	// Check if NIS already exists
	var existingSiswa Siswa
	if err := s.db.Where("nis = ?", req.NIS).First(&existingSiswa).Error; err == nil {
		return nil, errors.New("NIS sudah terdaftar")
	}

	// Check if email already exists (if provided)
	if req.Email != "" {
		if err := s.db.Where("email = ?", req.Email).First(&existingSiswa).Error; err == nil {
			return nil, errors.New("email sudah terdaftar")
		}
	}

	// Parse tanggal lahir
	var tanggalLahir *time.Time
	if req.TanggalLahir != "" {
		t, err := time.Parse("2006-01-02", req.TanggalLahir)
		if err != nil {
			return nil, errors.New("format tanggal lahir tidak valid, gunakan format YYYY-MM-DD")
		}
		tanggalLahir = &t
	}

	// Create siswa
	newSiswa := Siswa{
		NIS:          req.NIS,
		Nama:         req.Nama,
		JenisKelamin: req.JenisKelamin,
		TempatLahir:  req.TempatLahir,
		TanggalLahir: tanggalLahir,
		Alamat:       req.Alamat,
		NoTelepon:    req.NoTelepon,
		Email:        req.Email,
		Kelas:        req.Kelas,
		TahunMasuk:   req.TahunMasuk,
	}

	if err := s.db.Create(&newSiswa).Error; err != nil {
		return nil, err
	}

	return s.toResponseDTO(&newSiswa), nil
}

// GetAll retrieves all siswa with pagination and search.
func (s *service) GetAll(page, limit int, search string) (*dto.PaginatedResponse[SiswaResponseDTO], error) {
	var siswaList []Siswa
	var total int64

	// Default pagination
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	offset := (page - 1) * limit

	query := s.db.Model(&Siswa{})

	// Search by nama or NIS
	if search != "" {
		searchPattern := "%" + search + "%"
		query = query.Where("nama ILIKE ? OR nis ILIKE ?", searchPattern, searchPattern)
	}

	// Get total count
	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}

	// Get paginated data
	if err := query.Offset(offset).Limit(limit).Order("created_at DESC").Find(&siswaList).Error; err != nil {
		return nil, err
	}

	// Convert to response DTOs
	responseList := make([]SiswaResponseDTO, 0, len(siswaList))
	for _, siswa := range siswaList {
		responseList = append(responseList, *s.toResponseDTO(&siswa))
	}

	result := dto.NewPaginatedResponse(responseList, total, page, limit)
	return &result, nil
}

// GetByID retrieves a siswa by ID.
func (s *service) GetByID(id uint) (*SiswaResponseDTO, error) {
	var siswa Siswa
	if err := s.db.First(&siswa, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("siswa tidak ditemukan")
		}
		return nil, err
	}

	return s.toResponseDTO(&siswa), nil
}

// Update updates a siswa by ID.
func (s *service) Update(id uint, req UpdateSiswaRequestDTO) (*SiswaResponseDTO, error) {
	var siswa Siswa
	if err := s.db.First(&siswa, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("siswa tidak ditemukan")
		}
		return nil, err
	}

	// Check if NIS is being changed and already exists
	if req.NIS != "" && req.NIS != siswa.NIS {
		var existingSiswa Siswa
		if err := s.db.Where("nis = ? AND id != ?", req.NIS, id).First(&existingSiswa).Error; err == nil {
			return nil, errors.New("NIS sudah terdaftar")
		}
		siswa.NIS = req.NIS
	}

	// Check if email is being changed and already exists
	if req.Email != "" && req.Email != siswa.Email {
		var existingSiswa Siswa
		if err := s.db.Where("email = ? AND id != ?", req.Email, id).First(&existingSiswa).Error; err == nil {
			return nil, errors.New("email sudah terdaftar")
		}
		siswa.Email = req.Email
	}

	// Update fields if provided
	if req.Nama != "" {
		siswa.Nama = req.Nama
	}
	if req.JenisKelamin != "" {
		siswa.JenisKelamin = req.JenisKelamin
	}
	if req.TempatLahir != "" {
		siswa.TempatLahir = req.TempatLahir
	}
	if req.TanggalLahir != "" {
		t, err := time.Parse("2006-01-02", req.TanggalLahir)
		if err != nil {
			return nil, errors.New("format tanggal lahir tidak valid, gunakan format YYYY-MM-DD")
		}
		siswa.TanggalLahir = &t
	}
	if req.Alamat != "" {
		siswa.Alamat = req.Alamat
	}
	if req.NoTelepon != "" {
		siswa.NoTelepon = req.NoTelepon
	}
	if req.Kelas != "" {
		siswa.Kelas = req.Kelas
	}
	if req.TahunMasuk != 0 {
		siswa.TahunMasuk = req.TahunMasuk
	}

	if err := s.db.Save(&siswa).Error; err != nil {
		return nil, err
	}

	return s.toResponseDTO(&siswa), nil
}

// Delete soft deletes a siswa by ID.
func (s *service) Delete(id uint) error {
	result := s.db.Delete(&Siswa{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("siswa tidak ditemukan")
	}
	return nil
}

// toResponseDTO converts a Siswa model to SiswaResponseDTO.
func (s *service) toResponseDTO(siswa *Siswa) *SiswaResponseDTO {
	return &SiswaResponseDTO{
		ID:           siswa.ID,
		NIS:          siswa.NIS,
		Nama:         siswa.Nama,
		JenisKelamin: siswa.JenisKelamin,
		TempatLahir:  siswa.TempatLahir,
		TanggalLahir: siswa.TanggalLahir,
		Alamat:       siswa.Alamat,
		NoTelepon:    siswa.NoTelepon,
		Email:        siswa.Email,
		Kelas:        siswa.Kelas,
		TahunMasuk:   siswa.TahunMasuk,
		CreatedAt:    siswa.CreatedAt,
		UpdatedAt:    siswa.UpdatedAt,
	}
}
