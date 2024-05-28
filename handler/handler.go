package handler

import (
	"github.com/PedroCost4/movies-backend/config"
	"github.com/jmoiron/sqlx"
)

var (
	logger *config.Logger
	db     *sqlx.DB
)

func InitHandlers() {
	logger = config.GetLogger("user")
	db = config.GetDb()
}
