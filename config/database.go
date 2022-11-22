package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	entity "github.com/techlateef/jwt-Auth-techies/entities"

	"gorm.io/driver/mysql"

	// "gorm.io/driver/postgres"

	// "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabaseConnection() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load env file")

	}

	dbDatabase := os.Getenv("DB_DATABASE")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	fmt.Print(dbDatabase)
	fmt.Println(dbHost)
	fmt.Println(dbPort)
	// mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbDatabase)
	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai", dbHost, dbUser, dbPassword,  dbDatabase)

	// fmt.Print(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to load env file")

	}
	db.AutoMigrate(&entity.User{})
	return db

}

func CloseDatabasec(db *gorm.DB) {
	dbPostgre, err := db.DB()
	if err != nil {
		panic("Failed ")
	}
	dbPostgre.Close()
}
