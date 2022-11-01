package authorizes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/skp/app/models"
	"github.com/skp/app/services"
)

type AuthorizeUser struct {
	Ser *services.Services
}

func (authorizeUser *AuthorizeUser) SetAuth(ctx *fiber.Ctx) error {

	return ctx.Next()
}

func (authorizeUser *AuthorizeUser) CheckAuths(ctx *fiber.Ctx) error {
	cookie := ctx.Cookies("access_token")
	user := models.User{}
	accestoken := cookie
	// fmt.Println("Cookies ", accestoken)
	auths := authorizeUser.Ser.RedisClient.Get(accestoken, &user)
	// fmt.Println("Redis ", user)
	if auths != nil {
		// return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		// 	"status":  "FALIED",
		// 	"message": "Unauthorized",
		// 	"error":   "",
		// })
		return ctx.Render("unauthorized", fiber.Map{
			"Message": " Unauthorized",
			"Code":    fiber.StatusUnauthorized,
			"Title":   "SKP",
		})
	}
	return ctx.Next()
}
