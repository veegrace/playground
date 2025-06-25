package main

import (
	"fmt"
	"log"

	"github.com/microsoft/go-mssqldb/azuread"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"gorm.io/plugin/opentelemetry/tracing"
)

func SetupDBConnection() (*gorm.DB, error) {
	username := "sa"
	password := "PeroxideAce13!"
	host := "127.0.0.1"
	port := "1433"

	dsn := fmt.Sprintf("sqlserver://%v:%v@%v:%v", username, password, host, port)

	var (
		db  *gorm.DB
		err error
	)

	dialector := &sqlserver.Dialector{
		Config: &sqlserver.Config{
			DSN:        dsn,
			DriverName: azuread.DriverName,
		},
	}

	db, err = gorm.Open(dialector, &gorm.Config{
		Logger: gormlogger.Default.LogMode(gormlogger.Silent),
	})

	if err := db.Use(tracing.NewPlugin()); err != nil {
		panic(err)
	}

	if err != nil {
		panic(fmt.Errorf("gorm err: %v", err))
	}

	pool, err := db.DB()
	if err != nil {
		panic(fmt.Errorf("db err: %v", err))
	}

	pool.SetConnMaxLifetime(0)
	pool.SetMaxIdleConns(50)
	pool.SetMaxOpenConns(50)
	return db, nil
}

func main() {
	db, err := SetupDBConnection()
	if err != nil {
		panic(err)
	}

	db = db.Exec("CREATE DATABASE aespa;")
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	if err := sqlDB.Ping(); err != nil {
		panic(err)
	}
	log.Printf("DONE")
}
