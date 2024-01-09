package repository

import (
	"golangsidang/models"
	"log"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type ProgramRepository struct {
	DB *gorm.DB
}

func NewProgramRepository(db *gorm.DB) *ProgramRepository {
	return &ProgramRepository{DB: db}
}

func init() {
	err := godotenv.Load(filepath.Join(".", ".env"))
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}
}

func (r *ProgramRepository) CreateProgram(c *fiber.Ctx) error {
	programrequest := new(models.Program)
	if err := c.BodyParser(programrequest); err != nil {
		c.Status(fiber.StatusUnprocessableEntity).JSON(
			&fiber.Map{
				"status": "error",
			})
		return err
	}
	// 1-3 level
	if programrequest.Level < 1 || programrequest.Level > 3 {
		c.Status(fiber.StatusUnprocessableEntity).JSON(
			&fiber.Map{
				"status":  "error",
				"message": "Invalid level value",
			})
		return nil
	}

	programs := models.Program{
		Title:            programrequest.Title,
		Slug:             programrequest.Slug,
		Description:      programrequest.Description,
		Image_destop:     programrequest.Image_destop,
		Image_mobile:     programrequest.Image_mobile,
		Level:            programrequest.Level,
		Is_certification: programrequest.Is_certification,
		Url_Logo:         programrequest.Url_Logo,
		Pic_name:         programrequest.Pic_name,
		Pic_phone:        programrequest.Pic_phone,
		Start_at:         programrequest.Start_at,
		End_at:           programrequest.End_at,
		Is_active:        programrequest.Is_active,
		Is_publish:       programrequest.Is_publish,
		Create_at:        time.Now().Format("2006-01-02 15:04:05"),
		Delete:           false,
	}

	err := r.DB.Create(&programs).Error
	if err != nil {
		c.Status(fiber.StatusUnprocessableEntity).JSON(
			&fiber.Map{
				"status": "error",
			})
		return err
	}
	c.Status(fiber.StatusCreated).JSON(
		&fiber.Map{
			"status": "success",
			"data":   programs,
		})
	return nil
}
