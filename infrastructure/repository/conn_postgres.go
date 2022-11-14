package repository

import (
	"database/sql"
	"fmt"
	"tech-user-balance/config"
	"tech-user-balance/pkg"
)

func ConnectDB(conf *config.Config, log *pkg.Logger) (db *sql.DB) {
	psqlConnStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.ConfigDataBase.Host, conf.ConfigDataBase.Port, conf.ConfigDataBase.User, conf.ConfigDataBase.Password, conf.ConfigDataBase.NameDataBase)
	db, err := sql.Open("postgres", psqlConnStr)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return
}
