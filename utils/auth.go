package utils

import (
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

const jwtSecret = "your_dynamic_secret_key"

func TokenMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")

	if authHeader == "" {
		return c.Status(http.StatusUnauthorized).JSON(
			&fiber.Map{"message": "Unauthorized - Missing token"})
	}

	// Check for the "Bearer" prefix
	const bearerPrefix = "Bearer "
	if !strings.HasPrefix(authHeader, bearerPrefix) {
		return c.Status(http.StatusUnauthorized).JSON(
			&fiber.Map{"message": "Unauthorized - Invalid token format"})
	}

	// Extract the token without the "Bearer" prefix
	tokenString := authHeader[len(bearerPrefix):]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if err != nil || !token.Valid {
		return c.Status(http.StatusUnauthorized).JSON(
			&fiber.Map{"message": "Unauthorized - Invalid token"})
	}

	c.Locals("user", token.Claims.(jwt.MapClaims)["username"])
	c.Locals("password", token.Claims.(jwt.MapClaims)["password"])
	return c.Next()
}
