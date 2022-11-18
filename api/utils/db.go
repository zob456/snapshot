package utils

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

var (
	host     = os.Getenv("DB_HOST")
	port     = os.Getenv("DB_PORT")
	dbname   = os.Getenv("DB_NAME")
	user     = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	sslMode  = os.Getenv("SSL_MODE")
)

func ConnectDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslMode)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalln(err)
	}
	return db

}