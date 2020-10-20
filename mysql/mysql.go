package mysql

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql" // go get -u github.com/go-sql-driver/mysql
	m "github.com/ssoyyoung.p/seoulbitz-Backend/model"
	utils "github.com/ssoyyoung.p/seoulbitz-Backend/utils"
)

// get DB info func
func getDBinfo() string {
	data, _ := os.Open("mysql/info.json")

	var info m.DBinfo
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

// SingleDataQuery func
func SingleDataQuery() string {
	DB := ConnectDB()
	defer DB.Close()

	var title string
	query := "SELECT title FROM foodie LIMIT 1"
	err := DB.QueryRow(query).Scan(&title)

	utils.CheckErr(err)

	return title

}

// GetFoddieList func
func GetFoddieList() []m.Foddie {
	DB := ConnectDB()
	defer DB.Close()

	var foddie m.Foddie
	var allFoddie []m.Foddie

	query := "SELECT xpoint, ypoint, title, tag, like_cnt, addr, insta, thumb FROM `foodie`"
	rows, err := DB.Query(query)
	utils.CheckErr(err)

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&foddie.Xpoint, &foddie.Ypoint, &foddie.Title, &foddie.Tag, &foddie.LikeCnt, &foddie.Addr, &foddie.Insta, &foddie.Thumb)
		utils.CheckErr(err)

		allFoddie = append(allFoddie, foddie)
	}

	return allFoddie
}

//struct type to json
// 1. json.Marchal
// 2. string

// ex
// js, err := json.Marshal(foddie)
// utils.CheckErr(err)
// fmt.Println(string(js))
