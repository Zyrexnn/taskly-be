package user

// RegisterRequestDTO defines the structure for the user registration request body.
type RegisterRequestDTO struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

// LoginRequestDTO defines the structure for the user login request body.
type LoginRequestDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// UserResponseDTO defines the structure for user data in responses (without password).
type UserResponseDTO struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// LoginResponseDTO defines the structure for the login response, including the JWT.
type LoginResponseDTO struct {
	Token string          `json:"token"`
	User  UserResponseDTO `json:"user"`
}
