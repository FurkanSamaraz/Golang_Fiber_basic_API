package main

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	Ad     string `json:"ad"`
	Soyad  string `json:"soyad"`
	Numara int    `json:"numara"`
}

var user = []User{
	{Ad: "furkan", Soyad: "samaraz", Numara: 123},
	{Ad: "Sinan", Soyad: "celik", Numara: 321},
	{Ad: "sefa", Soyad: "pamuk", Numara: 435},
}

func tablo(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(user)
}
func getuser(c *fiber.Ctx) error {
	username := c.Params("ad")
	for _, element := range user {
		if username == element.Ad {
			return c.Status(http.StatusOK).JSON(element)
		}
	}

	return c.Status(http.StatusOK).JSON("ok")

}

func create(c *fiber.Ctx) error {
	var newCustomers User

	if err := c.BodyParser(&newCustomers); err != nil {
		return c.Status(http.StatusBadRequest).JSON("create")
	}
	user = append(user, newCustomers)

	return c.Status(http.StatusOK).JSON("Success ")

}
func Update(c *fiber.Ctx) error {
	var newCustomers User

	if err := c.BodyParser(&newCustomers); err != nil {
		return c.Status(http.StatusBadRequest).JSON("update")
	}

	for i, element := range user {
		if element.Ad == newCustomers.Ad {
			element.Soyad = newCustomers.Soyad
			element.Numara = newCustomers.Numara
			user = append(user[:i], user[i+1:]...)
			user = append(user, element)
		}
	}

	return c.Status(http.StatusOK).JSON("Updated Data")
}
func Delete(c *fiber.Ctx) error {

	username := c.Params("ad")
	for i, element := range user {
		if element.Ad == username {
			user = append(user[:i], user[i+1:]...)
		}
	}
	return c.Status(http.StatusOK).JSON("delete ok")
}
func main() {
	app := fiber.New()

	app.Get("/user/", tablo)
	app.Get("/user/:username", getuser)
	app.Post("/user/create", create)
	app.Post("/user/update", Update)

	app.Delete("/user/delete/:ad", Delete)

	app.Listen(":8080")
}
