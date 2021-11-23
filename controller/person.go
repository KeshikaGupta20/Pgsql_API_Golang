package controller

import (
	"fmt"

	"github.com/KeshikaGupta20/Pgsql_API_Golang/database"
	"github.com/KeshikaGupta20/Pgsql_API_Golang/models"

	"github.com/gofiber/fiber/v2"
)

func GetPerson(c *fiber.Ctx) error {

	ID := c.Params("id")

	db := database.DB

	var person []models.Person = make([]models.Person, 0)
	var books []models.Book

	db.First(&person, ID)
	db.Model(&person).Related(&books)

	//person.Books = books

	return c.JSON(fiber.Map{

		"status":  "success",
		"message": "Product found",
		"person":  person,
	})

}

func GetPeople(c *fiber.Ctx) error {

	db := database.DB

	var people []models.Person = make([]models.Person, 0)

	db.Find(&people)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{

		"success": true,
		"data": fiber.Map{
			"Persons": "people",
		},
	})

}

func CreatePerson(c *fiber.Ctx) error {

	db := database.DB

	p := new(models.Person)

	c.BodyParser(p)

	createdPerson := db.Create(&p)

	err := createdPerson.Error
	if err != nil {
		fmt.Println(err)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{

		"Message": "person added sucessfully",
	})

}

func DeletePerson(c *fiber.Ctx) error {

	ID := c.Params("id")

	db := database.DB

	var person []models.Person = make([]models.Person, 0)

	db.First(&person, ID)

	db.Delete(&person)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{

		"Message": "Product deleted sucessfully",
	})

}

func GetBook(c *fiber.Ctx) error {

	ID := c.Params("id")

	db := database.DB

	var book []models.Book

	db.First(&book, ID)

	return c.JSON(fiber.Map{

		"status":  "success",
		"message": "Product found",
		"Book":    book,
	})
}

func GetBooks(c *fiber.Ctx) error {

	var books []models.Book
	db := database.DB

	db.Find(&books)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{

		"success": true,
		"data": fiber.Map{
			"Books": "books",
		},
	})
}

func CreateBook(c *fiber.Ctx) error {

	db := database.DB

	b := new(models.Book)

	c.BodyParser(b)

	createdBook := db.Create(&b)

	err := createdBook.Error
	if err != nil {
		fmt.Println(err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{

		"Message": "book added sucessfully",
	})

}

func DeleteBook(c *fiber.Ctx) error {

	ID := c.Params("id")

	db := database.DB

	var book []models.Book

	db.First(&book, ID)
	db.Delete(&book)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{

		"Message": "Product deleted sucessfully",
	})
}
