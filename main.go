package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html"
	"github.com/skp/app/services"
	"github.com/skp/pkg/gormmanager"
	"github.com/skp/router"
)

func main() {
	gormmanager.New().GetInstanceConnect()
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/", "./assets")
	app.Use(recover.New())
	services := services.New()
	defer services.Db.Conn.Close()
	router.Routes(app, services)

	app.Listen(":3000")

}
