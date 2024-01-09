// main.go
package main

import (
	"log"
	"os"

	"golangsidang/models"
	"golangsidang/repository"
	"golangsidang/routes"
	"golangsidang/storage"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	// Setup database connection
	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASSWORD"),
		User:     os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	db, err := storage.NewConnection(config)
	if err != nil {
		log.Fatal("error connecting to the database")
	}

	// Migrate database models
	err = models.MigrateBooks(db)
	models.MigrateCourse(db)
	models.MigrateProgram(db)
	if err != nil {
		log.Fatal("error migrating books")
	}

	// Create a new Fiber app

	// Create a new Fiber app
	app := fiber.New()
	repo2 := repository.New(db)
	course := repository.CourseRepository{db}
	repo3 := repository.ProgramRepository{db}

	routes.UserRoutes(app, repo2)
	routes.CourseRoutes(app, &course)
	routes.ProgramRoutes(app, &repo3)

	// Create a repository instance using NewRepository
	repo := repository.NewRepository(db)
	// Create another repository instance using New
	// Setup routes with the first repository instance
	routes.SetupRoutes(app, repo)
	// Setup routes with the second repository instance
	// Start the Fiber app
	log.Fatal(app.Listen(":3000")) // Change the port if needed
}
