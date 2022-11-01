package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/skp/app/models"
	"github.com/skp/app/services"
)

type AdminHadler struct {
	Ser *services.Services
}

func (auth *AdminHadler) Home(ctx *fiber.Ctx) error {
	var dataCart []models.ResponseCart
	var quis models.Quis
	auth.Ser.Repository.GetResultQuisioner(&dataCart)
	err, getquis := auth.Ser.Repository.GetRowQuis(&quis)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"Error": err,
		})
	}
	return ctx.Render("admin", fiber.Map{"Title": "SKP", "dataCart": dataCart, "Question": getquis.Question})
}

func (auth *AdminHadler) Quis(ctx *fiber.Ctx) error {
	return ctx.Render("quis", fiber.Map{"Title": "SKP"})
}
func (auth *AdminHadler) CreateQuis(ctx *fiber.Ctx) error {
	reqquis := new(models.Quis)
	_ = ctx.BodyParser(reqquis)
	auth.Ser.Validator.ValidateRequest(reqquis)
	err := auth.Ser.Repository.CreateQuis(reqquis)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Message": "error"})
	}
	return ctx.RedirectBack("/createquis")
}
func (auth *AdminHadler) Listquesioner(ctx *fiber.Ctx) error {
	var quisioner []models.Quisioner
	auth.Ser.Repository.GetListQuisioner(&quisioner)
	// fmt.Println(quisioner)
	return ctx.Render("listquesioner", fiber.Map{
		"Title":        "SKP",
		"ListQusioner": quisioner,
	})
}
