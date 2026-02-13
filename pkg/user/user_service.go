package user

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Service interface {
	Register(req RegisterRequestDTO) (*UserResponseDTO, error)
	Login(req LoginRequestDTO) (*LoginResponseDTO, error)
}

type service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) Service {
	return &service{db: db}
}

// Register creates a new user.
func (s *service) Register(req RegisterRequestDTO) (*UserResponseDTO, error) {
	// Check if email already exists
	var existingUser User
	if err := s.db.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		return nil, errors.New("email already exists")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Create user
	newUser := User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	if err := s.db.Create(&newUser).Error; err != nil {
		return nil, err
	}

	response := &UserResponseDTO{
		ID:    newUser.ID,
		Name:  newUser.Name,
		Email: newUser.Email,
	}

	return response, nil
}

// Login authenticates a user and returns a JWT.
func (s *service) Login(req LoginRequestDTO) (*LoginResponseDTO, error) {
	var user User

	identifier := req.Identifier
	if identifier == "" {
		identifier = req.Email
	}

	if identifier == "" {
		return nil, errors.New("email or username is required")
	}

	// Search by email OR name
	if err := s.db.Where("email = ? OR name = ?", identifier, identifier).First(&user).Error; err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Compare password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Generate JWT
	token, err := generateJWT(user)
	if err != nil {
		return nil, err
	}

	response := &LoginResponseDTO{
		Token: token,
		User: UserResponseDTO{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
	}

	return response, nil
}

// generateJWT creates a new JWT for a given user.
func generateJWT(user User) (string, error) {
	jwtSecret := os.Getenv("JWT_SECRET")

	// Set claims
	claims := jwt.MapClaims{
		"sub":   strconv.Itoa(int(user.ID)),
		"email": user.Email,
		"iat":   time.Now().Unix(),
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return t, nil
}
