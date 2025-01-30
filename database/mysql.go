package database

import (
	"database/sql"
	"fmt"
	"telegram-chatgpt-bot.com/m/config"

	_ "github.com/go-sql-driver/mysql"
)

// ConnectDB initializes the database connection
func ConnectDB() (*sql.DB, error) {
	dsn := config.GetEnv("MYSQL_DSN")
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	fmt.Println("Database connected")
	return db, nil
}
