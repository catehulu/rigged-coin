package dbrepo

import (
	"github.com/catehulu/rigged-coin/internal/config"
	"github.com/catehulu/rigged-coin/internal/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoDBRepo struct {
	App *config.AppConfig
	DB  *mongo.Database
}

func NewMongoDBRepo(conn *mongo.Database, a *config.AppConfig) repository.DatabaseRepo {
	return &mongoDBRepo{
		App: a,
		DB:  conn,
	}
}
