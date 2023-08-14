package main

import (
	"log"
	"os"
	"server/api-go-test/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

func main() {

	// load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASS"), // no password set
		DB:       0,                       // use default DB
	})

	app := fiber.New()

	// route
	app.Get("/", routes.Hello)
	app.Get("/todos", routes.GetTodos(client))
	app.Post("/todos", routes.CreateTodo(client))
	app.Put("/todos/:id", routes.UpdateTodo(client))
	app.Delete("/todos/:id", routes.DeleteTodo(client))

	app.Listen(":3000")
}
