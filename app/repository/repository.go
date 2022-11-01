package repository

import (
	"github.com/skp/pkg/gormmanager"
	"github.com/skp/pkg/redisclient"
)

type Repository struct {
	Gormdb      *gormmanager.GormDB
	RedisClient *redisclient.RedisClient
}

func New(db *gormmanager.GormDB, rd *redisclient.RedisClient) *Repository {
	return &Repository{
		Gormdb:      db,
		RedisClient: rd,
	}
}
