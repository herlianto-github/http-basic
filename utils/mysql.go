package utils

import (
	"database/sql"
	"fmt"
	"httpBasic/configs"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB(config *configs.AppConfig) *sql.DB {

	connectionString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", config.Database.Username, config.Database.Password, config.Database.Address, config.Database.Port, config.Database.Name)

	db, err := sql.Open(config.Database.Driver, connectionString)

	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db

}

// func InitialMigration (db *sql.DB){

// }
