package main

import (
	"database/sql"
	"log"

	"github.com/anangfirmansyah/portfolio-cms-backend/cmd/api"
	"github.com/anangfirmansyah/portfolio-cms-backend/configs"
	"github.com/anangfirmansyah/portfolio-cms-backend/db"
	"github.com/go-sql-driver/mysql"
)

func main() {
	db, err := db.NewMYSQLStorage(mysql.Config{
		User:                 configs.Envs.DBUser,
		Passwd:               configs.Envs.DBPassword,
		Addr:                 configs.Envs.DBAddress,
		DBName:               configs.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if err != nil {
		log.Fatal(err)
	}

	initStorate(db)

	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorate(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB: Successfully connected")
}
