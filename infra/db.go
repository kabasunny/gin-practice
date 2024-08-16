package infra

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres" // GORMのPostgreSQLドライバ
	"gorm.io/gorm"            // GORM ORMライブラリ
)

func SetupDB() *gorm.DB {
	dsn := fmt.Sprintf( // データベース接続に必要なデータソース名（DSN）をフォーマットして作成
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Faild to connect database")
	}

	return db
}
