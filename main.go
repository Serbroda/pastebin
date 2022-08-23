package main

import (
	"embed"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"pastebin/database"
)

//go:embed frontend/build
var reactApp embed.FS

func main() {
	loadEnv()

	var dbName = os.Getenv("DB_NAME")

	database.Connect(database.ConnectionOptions{Name: dbName})

	app := fiber.New(fiber.Config{
		DisableKeepalive: true,
	})
	app.Use(cors.New())

	serveStatic(app)

	app.Get("/api/v1/hello", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("Hi there")
	})

	log.Fatal(app.Listen(":3001"))
}

func serveStatic(app *fiber.App) {
	app.Use("/", filesystem.New(filesystem.Config{
		Root:       http.FS(reactApp),
		PathPrefix: "frontend/build",
	}))
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
