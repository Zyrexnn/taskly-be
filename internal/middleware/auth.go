package middleware

import (
	"os"
	"strconv"
	"strings"
	"tasklybe/internal/dto"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// Protected is a middleware function to protect routes that require authentication.
func Protected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(dto.NewErrorResponse("Missing authorization header", nil))
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(dto.NewErrorResponse("Invalid authorization header format", nil))
		}

		tokenString := parts[1]
		jwtSecret := os.Getenv("JWT_SECRET")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.NewError(fiber.StatusUnauthorized, "Unexpected signing method")
			}
			return []byte(jwtSecret), nil
		})

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(dto.NewErrorResponse("Invalid or expired token", err.Error()))
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Store user information in context
			userIDStr := claims["sub"].(string)
			userID, _ := strconv.ParseUint(userIDStr, 10, 32)
			
			c.Locals("userId", uint(userID))
			c.Locals("email", claims["email"].(string))
			return c.Next()
		}

		return c.Status(fiber.StatusUnauthorized).JSON(dto.NewErrorResponse("Invalid token claims", nil))
	}
}
