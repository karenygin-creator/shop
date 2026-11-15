package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var products = []Product{
	{ID: 1, Name: "Apple", Price: 100},
	{ID: 2, Name: "Banana", Price: 200},
	{ID: 3, Name: "Orange", Price: 300},
}

func main() {
	app := fiber.New()
	app.Get("/", func(c fiber.Ctx) error {
		return c.SendFile("index.html")
	})
	app.Get("/products", func(c fiber.Ctx) error {
		return c.JSON(products)
	})
	app.Post("/products", func(c fiber.Ctx) error {
		var product Product
		err := c.Bind().Body(&product)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err,
			})
		}
		product.ID = len(products) + 1
		products = append(products, product)
		return c.Status(fiber.StatusCreated).JSON(product)
	})
	app.Get("/products/:id", func(c fiber.Ctx) error {
		id := fiber.Params[int](c, "id")
		for _, product := range products {
			if product.ID == id {
				return c.JSON(product)
			}
		}
		return c.Status(404).JSON(fiber.Map{
			"message": "Product not found"})
	})
	app.Delete("/products/:id", func(c fiber.Ctx) error {
		id := fiber.Params[int](c, "id")
		for i, product := range products {
			if product.ID == id {
				products = append(products[:i], products[i+1:]...)
				return c.JSON(fiber.Map{
					"message": "Product deleted",
				})
			}
		}
		return c.Status(404).JSON(fiber.Map{
			"message": "Product not found"})
	})
	app.Post("/products/:id/delete", func(c fiber.Ctx) error {
		id := fiber.Params[int](c, "id")
		for i, product := range products {
			if product.ID == id {
				products = append(products[:i], products[i+1:]...)
				return c.JSON(fiber.Map{
					"message": "Product deleted",
				})
			}
		}
		return c.Status(404).JSON(fiber.Map{
			"message": "Product not found"})
	})
	log.Fatal(app.Listen(":3000"))
}
