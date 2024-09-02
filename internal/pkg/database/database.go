package database

import (
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func newDialector(cfg Config) (gorm.Dialector, error) {
	switch cfg.Type {
	case "mysql":
		return mysql.Open(cfg.DSN), nil
	case "sqlite":
		return sqlite.Open(cfg.DSN), nil
	case "postgres":
		return postgres.Open(cfg.DSN), nil
	}

	return nil, errors.New("unknown database type")
}

func Connect(cfg Config) error {
	dial, err := newDialector(cfg)
	if err != nil {
		return err
	}

	db, err = gorm.Open(dial)
	if err != nil {
		return err
	}

	return nil
}

func Migrate(models ...interface{}) error {
	for _, model := range models {
		if err := db.AutoMigrate(model); err != nil {
			return err
		}
	}

	return nil
}
