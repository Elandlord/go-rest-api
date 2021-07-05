package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Connection string
	Host       string
	Port       string
	Username   string
	Password   string
	Name       string
	Charset    string
}

func GetConfig() *Config {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_DATABASE")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	return &Config{
		DB: &DBConfig{
			Connection: "mysql",
			Host:       dbHost,
			Port:       dbPort,
			Username:   dbUsername,
			Password:   dbPassword,
			Name:       dbName,
			Charset:    "utf8",
		},
	}
}

func DbConnect(config *Config) *sql.DB {
	dbConnection, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(goDockerDB)/%s", config.DB.Username, config.DB.Password, config.DB.Name))

	if err != nil {
		panic(err)
	}

	dbConnection.SetConnMaxLifetime(time.Minute * 3)
	dbConnection.SetMaxOpenConns(10)
	dbConnection.SetMaxIdleConns(10)
	return dbConnection
}
