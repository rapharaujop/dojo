package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	dsn := "root:root@tcp(127.0.0.1:3306)/dojo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(
		&Person{},
	)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Post("people/", func(c *fiber.Ctx) error {
		var p Person
		if err := c.BodyParser(&p); err != nil {
			return err
		}
		db.Create(&p)
		// return c.SendString("Hello, Pessoa 👋!")
		return c.SendString(fmt.Sprintf("Nome: %s", p.Name))
	})

	log.Fatal(app.Listen(":3000"))

	// fmt.Println("Hello Word!")

	// var nomevar string
	// nomevar = "Eu"
	// fmt.Println(nomevar)

	// nomevarnovo := "Eu de novo"
	// fmt.Println(nomevarnovo)

	// var person Person
	// person.Name = "Raphael"
	// person.Age = 38

	// fmt.Println(person)

	// personNovo := Person{
	// 	Name: "Marina",
	// 	Age:  34,
	// }

	//CRIA A PERSON NO BANCO COM DADOS PASSADO POR ELA
	// db.Create(&personNovo)

	// log.Println(personNovo)
}
