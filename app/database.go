package app

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/tkanos/gonfig"
)

// func New(db *sql.DB) *DB {
// 	return &DB{
// 		db: db,
// 	}
// }

type DB struct {
	db *sql.DB
}

type Configuration struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOST     string
	DB_PORT     int
	DB_NAME     string
}

func GetConfig() Configuration {
	conf := Configuration{}
	gonfig.GetConf("config.json", &conf)
	return conf
}

func Init() *sql.DB {
	conf := GetConfig()
	// dsn := fmt.Sprintf("%s:%s@/%s?parseTime=true", conf.DB_USERNAME, conf.DB_PASSWORD, conf.DB_NAME)
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		conf.DB_HOST, conf.DB_USERNAME, conf.DB_PASSWORD, conf.DB_PORT, conf.DB_NAME)

	// db, err := sql.Open("mysql", "root@tcp(localhost:3306)/nidzamUserTest")
	// helper.PanicIfError(err)

	// db.SetMaxIdleConns(5)
	// db.SetMaxOpenConns(20)
	// db.SetConnMaxLifetime(60 * time.Minute)
	// db.SetConnMaxIdleTime(10 * time.Minute)

	// return db

	db, err := sql.Open("mssql", connString)
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)
	return db
}
