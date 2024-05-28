package config

import (
	"github.com/jmoiron/sqlx"
)

var (
	logger *Logger
	db     *sqlx.DB
)

func Init() error {
	var err error

	logger = GetLogger("config")

	db, err = DatabaseInit()
	if err != nil {
		logger.Errorf("Unable to connect to database: %v\n", err)
		return err
	}

	return nil
}

func GetLogger(p string) *Logger {
	logger := NewLogger(p)
	return logger
}

func GetDb() *sqlx.DB {
	return db
}
