package repository

import (
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"golangsidang/models"
	"golangsidang/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Repositorry struct {
	DB *gorm.DB
}

func New(db *gorm.DB) *Repositorry {
	return &Repositorry{DB: db}
}

func init() {
	// Load variables from .env file
	err := godotenv.Load(filepath.Join(".", ".env"))
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}
}

func (r *Repositorry) CreateUser(c *fiber.Ctx) error {
	userRequest := new(models.User)
	if err := c.BodyParser(userRequest); err != nil {
		c.Status(fiber.StatusUnprocessableEntity).JSON(
			&fiber.Map{
				"status":  "error",
				"message": "Invalid JSON",
			})
		return err
	}

	// Hash the user's password before saving it to the database
	hashedPassword, err := utils.HashPassword(userRequest.Password)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(
			&fiber.Map{"message": "Error hashing password"})
		return err
	}

	var users models.User
	r.DB.Where("username = ?", userRequest.Username).First(&users)
	if users.Username != "" {
		c.Status(fiber.StatusBadRequest).JSON(
			&fiber.Map{
				"status":  "error",
				"message": "already exist",
			})
		return nil
	}

	r.DB.Where("email = ?", userRequest.Email).First(&users)
	if users.Email != "" {
		c.Status(fiber.StatusBadRequest).JSON(
			&fiber.Map{
				"status":  "error",
				"message": "already exist",
			})
		return nil
	}

	//harusnya menggunakna @gmail.com
	if !endsWithGmail(userRequest.Email) {
		userRequest.Email += "@gmail.com"
	}
	user := models.User{
		Username:  userRequest.Username,
		Password:  hashedPassword,
		Email:     userRequest.Email,
		Image:     userRequest.Image,
		Create_at: time.Now().Format(time.RFC3339),
		Create_by: userRequest.Username,
		Role:      "user",
		Delete_at: false,
	}

	err = r.DB.Create(&user).Error
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(
			&fiber.Map{"message": "Error creating user"})
		return err
	}

	c.Status(fiber.StatusCreated).JSON(
		&fiber.Map{
			"status": "success",
			"data":   user,
		})
	return nil
}
func (r *Repositorry) Login(c *fiber.Ctx) error {
	loginRequest := new(struct {
		Username string `json:"username"`
		Password string `json:"password"`
	})

	if err := c.BodyParser(loginRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			&fiber.Map{"message": "Invalid request payload"})
	}

	// Retrieve the user from the database
	var user models.User
	err := r.DB.Where("username = ?", loginRequest.Username).First(&user).Error

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(
			&fiber.Map{"message": "Invalid username or password"})
	}

	// Check the password
	if !utils.CheckPasswordHash(loginRequest.Password, user.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(
			&fiber.Map{"message": "Invalid username or password"})
	}

	// Generate JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["password"] = user.Password
	claims["username"] = user.Username
	claims["exp"] = jwt.TimeFunc().Add(time.Hour * 24).Unix() // Set token expiration time

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(
			&fiber.Map{"message": "JWT secret not set"})
	}

	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			&fiber.Map{"message": "Error generating JWT token"})
	}
	// Return the generated JWT token in the response
	return c.JSON(&fiber.Map{
		"token": tokenString,
	})
}
func endsWithGmail(email string) bool {
	return len(email) >= len("@gmail.com") && email[len(email)-len("@gmail.com"):] == "@gmail.com"
}

// GetAllUser paginates and retrieves all users
func (r *Repositorry) GetAllUser(c *fiber.Ctx) error {
	var users []models.User
	var count int64

	r.DB.Model(&models.User{}).Count(&count)
	r.DB.Find(&users)

	// Get pagination parameters from query parameters
	pageNumber := c.Query("page", "1")
	pageSize := c.Query("pageSize", "10")

	// Convert string parameters to integers
	pageNum, err := strconv.Atoi(pageNumber)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid page number",
		})
	}

	pageSze, err := strconv.Atoi(pageSize)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid page size",
		})
	}

	// Paginate the users
	offset := (pageNum - 1) * pageSze
	r.DB.Offset(offset).Limit(pageSze).Find(&users)

	totalPages, pageSizeNow := utils.Pagination(count, pageNum, pageSze)

	return c.JSON(fiber.Map{
		"data": fiber.Map{
			"users": users,
			"meta": fiber.Map{
				"page":       pageNum,
				"pageSize":   pageSze,
				"totalPages": totalPages,
				"totalItems": pageSizeNow,
			},
		},
		"status": "success",
	})
}
func (r *Repositorry) GetUserByUsername(c *fiber.Ctx) error {
	username := c.Params("username")

	var user models.User
	r.DB.Where("username = ?", username).First(&user)

	return c.JSON(&fiber.Map{
		"status": "success",
		"data":   user,
	})
}

func (r *Repositorry) GetUserById(c *fiber.Ctx) error {
	id := c.Params("id")

	var user models.User
	r.DB.Where("id = ?", id).First(&user)

	return c.JSON(&fiber.Map{
		"status": "success",
		"data":   user,
	})
}

func (r *Repositorry) UpdateUserByid(c *fiber.Ctx) error {
	id := c.Params("id")

	userRequest := new(models.User)
	if err := c.BodyParser(userRequest); err != nil {
		c.Status(fiber.StatusUnprocessableEntity).JSON(
			&fiber.Map{
				"status":  "error",
				"message": "Invalid JSON",
			})
		return err
	}

	var user models.User
	r.DB.Where("id = ?", id).First(&user)

	user.Username = userRequest.Username
	user.Email = userRequest.Email
	user.Image = userRequest.Image
	user.Update_by = userRequest.Username

	err := r.DB.Save(&user).Error
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(
			&fiber.Map{"message": "Error updating user by id"})
		return err
	}

	return c.JSON(&fiber.Map{
		"status": "success",
		"data":   user,
	})
}

func (r *Repositorry) ResetPaswword(c *fiber.Ctx) error {
	username := c.Params("username")

	userRequest := new(models.User)
	if err := c.BodyParser(userRequest); err != nil {
		c.Status(fiber.StatusUnprocessableEntity).JSON(
			&fiber.Map{
				"status":  "error",
				"message": "Invalid JSON",
			})
		return err
	}

	var user models.User
	r.DB.Where("username = ?", username).First(&user)

	hashedPassword, err := utils.HashPassword(userRequest.Password)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(
			&fiber.Map{"message": "Error hashing password"})
		return err
	}

	user.Password = hashedPassword
	user.Update_by = userRequest.Username

	err = r.DB.Save(&user).Error
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(
			&fiber.Map{"message": "Error updating user"})
		return err
	}

	return c.JSON(&fiber.Map{
		"status": "success",
		"data":   user,
	})
}

func (r *Repositorry) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	var user models.User
	r.DB.Where("id = ?", id).First(&user)

	user.Delete_at = true
	user.Delete_by = user.Username

	err := r.DB.Save(&user).Error
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(
			&fiber.Map{"message": "Error updating user"})
		return err
	}

	return c.JSON(&fiber.Map{
		"status": "success",
		"data":   user,
	})
}
