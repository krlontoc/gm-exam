package config

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
var dbError error
var locTZ *time.Location

type QueryConfig struct {
	Columns []string
	Limit   int
	Offset  int
	Order   string
	OrderBy string
}

func Init() {
	initDB()
	locTZ, _ = time.LoadLocation("Asia/Manila")
}

func initDB() {
	DB, dbError = gorm.Open(sqlite.Open(" gme.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if dbError != nil {
		panic("failed to connect database")
	}
}

func GetLocTZ() *time.Location {
	return locTZ
}

func SetQueryConfig(tx *gorm.DB, qc QueryConfig) *gorm.DB {
	if qc.Limit > 0 {
		tx.Limit(qc.Limit)
	}

	if qc.Offset > 0 {
		tx.Offset(qc.Offset)
	}

	if qc.OrderBy != "" {
		if qc.Order != "" {
			tx.Order(qc.OrderBy + " " + qc.Order)
		}

		tx.Order(qc.OrderBy)
	}

	return tx
}
