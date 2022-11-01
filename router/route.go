package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/skp/app/handlers"
	"github.com/skp/app/services"
	"github.com/skp/pkg/authorizes"
)

func Routes(app *fiber.App, services *services.Services) {

	route := app.Group("/")
	fontend := &handlers.SKP{
		Ser: services,
	}
	login := &handlers.AuthHandler{
		Ser: services,
	}
	route.Get("", fontend.Home)
	route.Post("/create", fontend.CreateSKP)
	route.Get("/login", login.Login)
	route.Post("/login", login.Auth)
	auten := &authorizes.AuthorizeUser{
		Ser: services,
	}
	routeAuth := app.Group("/", auten.CheckAuths)
	Admin := &handlers.AdminHadler{
		Ser: services,
	}
	routeAuth.Get("/admin", Admin.Home)
	routeAuth.Get("/logout", login.Logout)
	routeAuth.Get("/quis", Admin.Quis)
	routeAuth.Post("/createquis", Admin.CreateQuis)
	routeAuth.Get("/listquesioner", Admin.Listquesioner)

}
