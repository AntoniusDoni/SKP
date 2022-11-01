package services

import (
	"github.com/skp/app/repository"
	"github.com/skp/pkg/customeauth"
	"github.com/skp/pkg/gormmanager"
	"github.com/skp/pkg/redisclient"
	"github.com/skp/pkg/validator"
)

type Services struct {
	Db          *gormmanager.GormDB
	RedisClient *redisclient.RedisClient
	Validator   *validator.Validator
	Repository  *repository.Repository
	Auth        *customeauth.Customeauth
}

func New() *Services {

	validator := validator.New()
	gormdb := gormmanager.New()
	redisclient := redisclient.New()
	Auth := customeauth.New()
	repository := repository.New(gormdb, redisclient)

	return &Services{
		Db:          gormdb,
		RedisClient: redisclient,
		Validator:   validator,
		Repository:  repository,
		Auth:        Auth,
	}
}
