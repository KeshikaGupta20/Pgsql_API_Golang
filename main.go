package main

import (
	"log"

	"github.com/KeshikaGupta20/Pgsql_API_Golang/database"
	"github.com/KeshikaGupta20/Pgsql_API_Golang/router"
	"github.com/joho/godotenv"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/v2/middleware/logger"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {

	app := fiber.New()

	app.Use(logger.New())

	godotenv.Load()

	database.ConnectDB()

	router.RegisterRoutes(app)

	log.Fatal(app.Listen(":3000"))

}
