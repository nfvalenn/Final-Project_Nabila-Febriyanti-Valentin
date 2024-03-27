package lib

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func InitDatabase() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=12345 dbname=FP port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}
	DB = db
	return DB, nil
}
