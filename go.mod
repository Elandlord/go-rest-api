module mentechmedia.nl/rest-api

go 1.16

require (
	github.com/go-sql-driver/mysql v1.6.0
	github.com/gorilla/mux v1.8.0
	github.com/joho/godotenv v1.3.0
	mentechmedia.nl/config v0.0.0-00010101000000-000000000000 // indirect
	mentechmedia.nl/model v0.0.0-00010101000000-000000000000 // indirect
)

replace mentechmedia.nl/article => ../article

replace mentechmedia.nl/config => ./config

replace mentechmedia.nl/model => ./model
