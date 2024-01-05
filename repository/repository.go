package repository

import (
	"log"
	"net/http"
	"path/filepath"
	"time"

	"golangsidang/models"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{DB: db}
}
func init() {
	// Load variables from .env file
	err := godotenv.Load(filepath.Join(".", ".env"))
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}
}

func (r *Repository) CreateBooks(c *fiber.Ctx) error {
	bookrequest := new(models.BookRequest)
	if err := c.BodyParser(bookrequest); err != nil {
		c.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{
				"status": "error",
			})
		return err
	}
	var books models.Book
	r.DB.Where("author = ?", bookrequest.Author).First(&books)
	if books.Author != "" {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{
				"status":  "error",
				"message": "already exist",
			})
		return nil
	}
	if books.Author != "" || books.Title != "" {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{
				"status":  "error",
				"message": "already exist",
			})
		return nil
	}

	book := models.Book{
		Author:    bookrequest.Author,
		Title:     bookrequest.Title,
		Publisher: bookrequest.Publisher,
		Delete_at: false,
		Create_at: time.Now().Format(time.RFC3339),
		Create_by: bookrequest.Author, // Menggunakan nilai dari input "author" untuk Create_by juga
	}

	// Lakukan langkah-langkah berikut sesuai kebutuhan Anda (misalnya, simpan ke database, dll.)

	err := c.BodyParser(&book)

	if err != nil {
		c.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{
				"status": "error",
			})
		return err
	}

	err = r.DB.Create(&book).Error
	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "error creating book"})
		return err
	}

	c.Status(http.StatusBadRequest).JSON(
		&fiber.Map{
			"status": "success",
			"data":   book,
		})
	return nil
}

func (r *Repository) DeleteBook(context *fiber.Ctx) error {
	id := context.Params("id")

	if id == "" {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "id is required"})
		return nil
	}

	// Retrieve the book from the database
	var bookModel models.Book

	err := r.DB.Where("id = ?", id).First(&bookModel).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "error retrieving book"})
		return err
	}

	// Update the Delete_at field to true
	bookModel.Delete_at = true
	bookModel.Delete_by = bookModel.Author

	// Save the updated book to the database with a WHERE condition
	err = r.DB.Where("id = ?", id).Save(&bookModel).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "error updating book"})
		return err
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{
			"status": "success",
			"data":   bookModel,
		})
	return nil
}

func (r *Repository) GetBooks(c *fiber.Ctx) error {
	bookModels := &[]models.Books{}
	err := r.DB.Find(bookModels).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "error getting books"})
		return err
	}
	c.Status(http.StatusOK).JSON(
		&fiber.Map{
			"status": "success",
			"data":   bookModels,
		})
	return nil
}

func (r *Repository) UpdateBook(context *fiber.Ctx) error {
	id := context.Params("id")

	if id == "" {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "id is required"})
		return nil
	}

	// Retrieve the book from the database
	var bookModel models.Book

	err := r.DB.Where("id = ?", id).First(&bookModel).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "error retrieving book"})
		return err
	}

	// Update the book model
	bookrequest := new(models.BookRequest)
	if err := context.BodyParser(bookrequest); err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{
				"status": "error",
			})
		return err
	}
	bookModel.Author = bookrequest.Author
	bookModel.Title = bookrequest.Title
	bookModel.Publisher = bookrequest.Publisher
	// Update the Delete_at field to true
	bookModel.Update_by = true
	bookModel.Update_at = time.Now().Format(time.RFC3339)

	// Save the updated book to the database with a WHERE condition
	err = r.DB.Where("id = ?", id).Save(&bookModel).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "error updating book"})
		return err
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{
			"status": "success",
			"data":   bookModel,
		})
	return nil
}

func (r *Repository) GetBook(c *fiber.Ctx) error {
	book := &models.Books{}
	id := c.Params("id")
	if id == "" {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "id is required"})
		return nil
	}
	err := r.DB.First(book, id).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "error getting book"})
		return err
	}
	c.Status(http.StatusOK).JSON(
		&fiber.Map{
			"status": "success",
			"data":   book,
		})
	return nil
}
