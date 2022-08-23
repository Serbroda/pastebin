package main

import (
	"embed"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/joho/godotenv"
	"github.com/teris-io/shortid"
	"log"
	"net/http"
	"os"
	"pastebin/database"
	"pastebin/handlers"
	"pastebin/middlewares"
)

//go:embed frontend/build
var reactApp embed.FS

func main() {
	loadEnv()

	sid, _ := shortid.New(1, shortid.DefaultABC, 2342)
	shortid.SetDefault(sid)

	var dbName = os.Getenv("DB_NAME")

	database.Connect(database.ConnectionOptions{Name: dbName})

	app := fiber.New(fiber.Config{
		DisableKeepalive: true,
	})
	app.Use(cors.New())

	var store = session.New()
	app.Use(middlewares.NewSessionize(store))

	serveStatic(app)
	setupRoutes(app, store)

	log.Fatal(app.Listen(":3001"))
}

func serveStatic(app *fiber.App) {
	app.Use("/", filesystem.New(filesystem.Config{
		Root:       http.FS(reactApp),
		PathPrefix: "frontend/build",
	}))
}

func setupRoutes(app *fiber.App, store *session.Store) {
	v1 := app.Group("/api/v1")
	v1.Get("hello", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("Hi there")
	})

	pastes := v1.Group("/pastes")
	pastes.Post("/", handlers.CreatePaste(store))
	pastes.Get("/", handlers.GetPastes(store))
	pastes.Get("/:pasteId", handlers.GetPaste)
	pastes.Patch("/:pasteId", handlers.UpdatePaste)
	pastes.Delete("/:pasteId", handlers.DeletePaste)
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
