package mysql

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql" // go get -u github.com/go-sql-driver/mysql
	model "github.com/ssoyyoung.p/seoulbitz-Backend/model"
)

// get DB info func
func getDBinfo() string {
	data, _ := os.Open("mysql/info.json")

	var info model.DBinfo
	byteValue, _ := ioutil.ReadAll(data)
	json.Unmarshal(byteValue, &info)

	DBinfo := info.Username + ":" + info.Password + "@tcp(" + info.Hostname + info.Port + ")/" + info.Database

	return DBinfo

}

// ConnectDB func
func ConnectDB() *sql.DB {
	DBinfo := getDBinfo()

	db, err := sql.Open("mysql", DBinfo)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}
