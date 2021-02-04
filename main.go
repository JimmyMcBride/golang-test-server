package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

// SomeStruct is some strucutre that serves testing and exporational purposes.
type SomeStruct struct {
	Name string `json:"name"`
	Age  uint8  `json:"age"`
}

// SomeOtherStruct is some other strucutre that serves testing and exporational purposes.
type SomeOtherStruct struct {
	Count int    `json:"count"`
	Time  string `json:"time"`
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		data := SomeStruct{
			Name: "Grame",
			Age:  20,
		}

		return c.JSON(data)
	})

	app.Get("/loop", func(c *fiber.Ctx) error { // Time: 201 ms - Size: 122 B
		start := time.Now()
		var count int = 0
		for i := 0; i < 1000; i++ {
			count++
		}
		elapsed := time.Since(start)
		message := fmt.Sprintf("Time elapsed: %s", elapsed)
		data := SomeOtherStruct{
			Count: count,
			Time:  message,
		}

		return c.JSON(data)
	})

	log.Fatal(app.Listen(":3000"))
}
