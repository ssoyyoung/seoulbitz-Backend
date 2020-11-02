package mysql

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
	"strings"

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
		fmt.Println(err)
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
	query := "SELECT title FROM Foodie LIMIT 1"
	err := DB.QueryRow(query).Scan(&title)

	utils.CheckErr(err)

	return title

}

// GetFoodieList func
func GetFoodieList() []m.Foodie {
	DB := ConnectDB()
	defer DB.Close()

	var Foodie m.Foodie
	var allFoodie []m.Foodie

	query := "SELECT xpoint, ypoint, title, tag, like_cnt, addr, insta, thumb FROM `foodie`"
	rows, err := DB.Query(query)
	utils.CheckErr(err)

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&Foodie.Xpoint, &Foodie.Ypoint, &Foodie.Title, &Foodie.Tag, &Foodie.LikeCnt, &Foodie.Addr, &Foodie.Insta, &Foodie.Thumb)
		utils.CheckErr(err)

		allFoodie = append(allFoodie, Foodie)
	}

	return allFoodie
}

// GetShoppingList func
func GetShoppingList() []m.Shopping {
	DB := ConnectDB()
	defer DB.Close()

	var shop m.Shopping
	var allShops []m.Shopping

	query := "SELECT xpoint, ypoint, title, tag, like_cnt, addr, insta, thumb FROM `shopping`"
	rows, err := DB.Query(query)
	utils.CheckErr(err)

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&shop.Xpoint, &shop.Ypoint, &shop.Title, &shop.Tag, &shop.LikeCnt, &shop.Addr, &shop.Insta, &shop.Thumb)
		utils.CheckErr(err)

		allShops = append(allShops, shop)
	}

	return allShops
}

// GetSubwayList func
func GetSubwayList() []m.Subway {
	DB := ConnectDB()
	defer DB.Close()

	var subway m.Subway
	var allSubways []m.Subway

	query := "SELECT xpoint_wgs, ypoint_wgs, station_nm, station_cd, line_num, fr_code, cyber_st_code FROM `subway`"
	rows, err := DB.Query(query)
	utils.CheckErr(err)

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&subway.XpointWgs, &subway.YpointWgs, &subway.StationNm, &subway.StationCd, &subway.LineNum, &subway.FrCode, &subway.CyberStCode)
		utils.CheckErr(err)

		allSubways = append(allSubways, subway)
	}

	return allSubways
}

// GetSubwayLatLng func
func GetSubwayLatLng(subwayName string) m.Subway {
	DB := ConnectDB()
	defer DB.Close()

	var subway m.Subway

	query := "SELECT xpoint_wgs, ypoint_wgs, station_nm, station_cd, line_num, fr_code, cyber_st_code FROM `subway` WHERE station_nm=?"
	err := DB.QueryRow(query, subwayName).Scan(&subway.XpointWgs, &subway.YpointWgs, &subway.StationNm, &subway.StationCd, &subway.LineNum, &subway.FrCode, &subway.CyberStCode)
	utils.CheckErr(err)
	return subway
}

// GetPlaceLatLng func
func GetPlaceLatLng(dbName string) []m.PlaceLatLng {
	DB := ConnectDB()
	defer DB.Close()

	var place m.PlaceLatLng
	var allPlaces []m.PlaceLatLng

	query := "SELECT idx, title, xpoint, ypoint FROM " + dbName
	rows, err := DB.Query(query)
	utils.CheckErr(err)

	for rows.Next() {
		err := rows.Scan(&place.Idx, &place.Title, &place.XpointWgs, &place.YpointWgs)
		utils.CheckErr(err)

		allPlaces = append(allPlaces, place)
	}

	return allPlaces
}

// InsertFoodie func
func InsertFoodie(Foodie *m.Foodie) string {
	DB := ConnectDB()
	defer DB.Close()

	query := "INSERT INTO Foodie (xpoint, ypoint, title, tag, like_cnt, addr, insta, thumb) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	result, err := DB.Exec(query, Foodie.Xpoint, Foodie.Ypoint, Foodie.Title, Foodie.Tag, Foodie.LikeCnt, Foodie.Addr, Foodie.Insta, Foodie.Thumb)
	utils.CheckErr(err)

	n, err := result.RowsAffected()
	if n == 1 {
		fmt.Println("1 row inserted!")
	}

	return "Insert!"
}

// InsertShop func
func InsertShop(shopping *m.Shopping) string {
	DB := ConnectDB()
	defer DB.Close()

	query := "INSERT INTO shopping (xpoint, ypoint, title, tag, like_cnt, addr, insta, thumb) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	result, err := DB.Exec(query, shopping.Xpoint, shopping.Ypoint, shopping.Title, shopping.Tag, shopping.LikeCnt, shopping.Addr, shopping.Insta, shopping.Thumb)
	utils.CheckErr(err)

	n, err := result.RowsAffected()
	if n == 1 {
		fmt.Println("1 row inserted!")
	}

	return "Insert!"
}

// GetInfos func
func GetInfos(dbName string, places []string) []m.Foodie {
	DB := ConnectDB()
	defer DB.Close()

	placeList := strings.Join(places, "','")

	var Foodie m.Foodie
	var allFoodie []m.Foodie

	query := "SELECT xpoint, ypoint, title, tag, like_cnt, addr, insta, thumb FROM "+ dbName +" WHERE title IN ('"+placeList+"')"
	fmt.Println(query)
	rows, err := DB.Query(query)
	utils.CheckErr(err)

	for rows.Next() {
		err := rows.Scan(&Foodie.Xpoint, &Foodie.Ypoint, &Foodie.Title, &Foodie.Tag,
						&Foodie.LikeCnt, &Foodie.Addr, &Foodie.Insta, &Foodie.Thumb)
		utils.CheckErr(err)

		allFoodie = append(allFoodie, Foodie)
	}

	return allFoodie
}
