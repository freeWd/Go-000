package models2

import (
	"context"
	"database/sql"
	"fmt"
	"go-gin-app/pkg/setting"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	ctx context.Context
	db  *sql.DB
)

func init() {
	var (
		dbType, dbName, user, password, host string
	)
	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}

	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName)

	db, err = sql.Open(dbType, dsn)

	if err != nil {
		log.Println(err)
	}

	db.SetConnMaxLifetime(100)
	db.SetMaxIdleConns(10)
}
