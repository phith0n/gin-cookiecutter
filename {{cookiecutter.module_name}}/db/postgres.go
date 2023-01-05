package db

import (
	"{{cookiecutter.module_name}}/logging"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	dbLogger "gorm.io/gorm/logger"
	"time"
)

var logger = logging.GetSugar()

type Database struct {
	*gorm.DB
}

var DB *Database

func InitPostgres(databaseURL string, debug bool) error {
	gormConfig := &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	}
	if debug {
		gormConfig.Logger = &CustomLogger{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			BaseLevel:                 dbLogger.Info, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
		}
	}

	pg := postgres.Open(databaseURL)
	db, err := gorm.Open(pg, gormConfig)
	if err != nil {
		logger.Errorf("fail to connect to database: %v", err)
		return err
	}

	DB = &Database{
		DB: db,
	}
	return nil
}
