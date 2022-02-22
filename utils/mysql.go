package utils

import (
	"database/sql"
	"fmt"
	"httpBasic/configs"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func InitDB(config *configs.AppConfig) *sql.DB {

	connectionString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?multiStatements=true", config.Database.Username, config.Database.Password, config.Database.Address, config.Database.Port, config.Database.Name)

	db, err := sql.Open(config.Database.Driver, connectionString)
	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		"file://migrations",
		config.Database.Driver,
		driver,
	)
	m.Steps(2)

	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db

}
