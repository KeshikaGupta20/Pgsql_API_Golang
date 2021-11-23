package router

import (
	c "github.com/KeshikaGupta20/Pgsql_API_Golang/controller"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app fiber.Router) {

	app.Post("/createperson", c.CreatePerson)

	app.Post("/createbook", c.CreateBook)

	app.Delete("/deleteperson/{id}", c.DeletePerson)

	app.Delete("/deletebook/{id}", c.DeleteBook)

	app.Get("/getbooks", c.GetBooks)

	app.Get("/getbook/{id}", c.GetBook)

	app.Get("/getpeople", c.GetPeople)

	app.Get("/getperson/{id}",c. GetPerson)

}
