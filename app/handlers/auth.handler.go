package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/skp/app/models"
	"github.com/skp/app/services"
)

type AuthHandler struct {
	Ser *services.Services
}

func (auth *AuthHandler) Login(ctx *fiber.Ctx) error {
	return ctx.Render("login", fiber.Map{})
}
func (auth *AuthHandler) Auth(ctx *fiber.Ctx) error {
	userRequest := new(models.UserLoginReq)
	_ = ctx.BodyParser(userRequest)
	user := &models.User{}
	err := auth.Ser.Repository.GetUserEmailPass(userRequest.Usermail, user)
	if err != nil {
		// return ctx.Status(fiber.StatusExpectationFailed).JSON(fiber.Map{
		// 	"status":  "Failed",
		// 	"message": "Your Account not Exist",
		// 	"error":   "",
		// })
		return ctx.Render("unauthorized", fiber.Map{"Message": " Your Account not Exist", "Code": fiber.StatusExpectationFailed, "Title": "SKP"})
	}
	accessToken, _ := auth.Ser.Auth.CreateUserToken(userRequest.Usermail, user)
	cookie := fiber.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
		Secure:   true,
	}
	auth.Ser.RedisClient.Set(accessToken, user, time.Hour*24)
	ctx.Cookie(&cookie)
	return ctx.Redirect("/admin")
}

func (auth *AuthHandler) Logout(ctx *fiber.Ctx) error {
	cookie := ctx.Cookies("access_token")
	auth.Ser.RedisClient.Remove(cookie)
	return ctx.RedirectBack("/")
}
