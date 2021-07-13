package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"regexp"
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

const projectDirName = "rest-api"

func loadEnv() {
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	err := godotenv.Load(string(rootPath) + `/.env`)

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func loadTestingEnv() {
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	err := godotenv.Load(string(rootPath) + `/.env.testing`)

	if err != nil {
		log.Fatalf("Error loading .env.testing file")
	}
}

func GetConfig() *Config {
	loadEnv()

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

func GetTestingConfig() *Config {
	loadTestingEnv()

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
	dbConnection, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", config.DB.Username, config.DB.Password, config.DB.Host, config.DB.Name))

	if err != nil {
		panic(err)
	}

	dbConnection.SetConnMaxLifetime(time.Minute * 3)
	dbConnection.SetMaxOpenConns(10)
	dbConnection.SetMaxIdleConns(10)
	return dbConnection
}
