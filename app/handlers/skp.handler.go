package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/skp/app/models"
	"github.com/skp/app/services"
)

type SKP struct {
	Ser *services.Services
}

func (skps *SKP) Home(ctx *fiber.Ctx) error {
	var quis models.Quis
	err, getquis := skps.Ser.Repository.GetRowQuis(&quis)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"Error": err,
		})
	}

	return ctx.Render("home", fiber.Map{"Question": getquis.Question, "idquis": getquis.IdQuis})
}

func (skp *SKP) CreateSKP(ctx *fiber.Ctx) error {
	reqquisoner := new(models.ReqQuisioner)
	_ = ctx.BodyParser(reqquisoner)

	skp.Ser.Validator.ValidateRequest(reqquisoner)
	reqquisoner.ID = uuid.New()
	// fmt.Println(reqquisoner.ResultQuist)
	err := skp.Ser.Repository.CreateSKP(reqquisoner)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Message": "error"})
	}
	return ctx.RedirectBack("/")
}
