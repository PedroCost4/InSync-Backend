package config

import (
	"os"

	"github.com/PedroCost4/movies-backend/schema"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

func DatabaseInit() (*sqlx.DB, error) {
	logger = GetLogger("database")

	db, err := sqlx.Connect("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		logger.Errorf("Unable to connect to database: %v\n", err)
		return nil, err
	}

	userSchemas := schema.GetSchemasString()

	result, err := db.Exec(userSchemas)

	if err != nil {
		logger.Errorf("Unable to create schemas: %v\n", err)
		return nil, err
	}

	logger.Infof("Schemas created: %v\n", result)

	logger.Infof("Connected to database")

	return db, nil
}
