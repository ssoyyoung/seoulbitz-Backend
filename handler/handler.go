package handler

import (
	"fmt"
	"strconv"
	"strings"

	//"encoding/json"
	"net/http"

	"github.com/labstack/echo"
	m "github.com/ssoyyoung.p/seoulbitz-Backend/model"
	mysql "github.com/ssoyyoung.p/seoulbitz-Backend/mysql"
	utils "github.com/ssoyyoung.p/seoulbitz-Backend/utils"
)

// GetFoodieList func
func GetFoodieList(c echo.Context) error {
	FoodieList := mysql.GetFoodieList()

	for idx, foodie := range FoodieList {
		uniq := strings.Split(foodie.Insta, "/")[4]
		FoodieList[idx].Uniq = uniq
	}

	fmt.Printf("Total Foodie's list count is %d \n", len(FoodieList))

	return c.JSON(http.StatusOK, FoodieList)
}

// GetShoppingList func
func GetShoppingList(c echo.Context) error {
	shoppingList := mysql.GetShoppingList()

	for idx, shop := range shoppingList {
		uniq := strings.Split(shop.Insta, "/")[4]
		shoppingList[idx].Uniq = uniq
	}

	fmt.Printf("Total shopping's list count is %d \n", len(shoppingList))

	return c.JSON(http.StatusOK, shoppingList)
}

// GetPlaceList func
func GetPlaceList(c echo.Context) error {
	placeList := mysql.GetPlaceList()
	page := c.Param("page")

	var p int

	for idx, place := range placeList {
		uniq := strings.Split(place.Insta, "/")[4]
		placeList[idx].Uniq = uniq
	}

	if page != "" {
		p, _ = strconv.Atoi(page)
		placeList = placeList[20*p : 20*(p+1)]
	}

	fmt.Printf("Total place's count is %d \n", len(placeList)) // fmt.Printf("%d") : 정수형
	return c.JSON(http.StatusOK, placeList)
}

// GetSubwayList func
func GetSubwayList(c echo.Context) error {
	subwayList := mysql.GetSubwayList()

	fmt.Printf("Total subway's list count is %d \n", len(subwayList))

	return c.JSON(http.StatusOK, subwayList)
}

// GetNearFoodiePlace func
func GetNearFoodiePlace(c echo.Context) error {
	subwayName := c.FormValue("subway")
	shopType := "foodie"

	subWayLatLng := mysql.GetSubwayLatLng(subwayName)
	if subWayLatLng.StationNm == "" {
		return c.String(http.StatusOK, "올바른 지하철역의 이름을 다시 입력해주세요")
	}
	placeList := mysql.GetPlaceLatLng(shopType)

	AllPointDis := utils.CalculateDistance(subwayName, subWayLatLng, placeList)

	var resultPlace []string
	kvale := map[string]float64{}

	for _, Point := range AllPointDis[:10] {
		resultPlace = append(resultPlace, Point.Title)
		kvale[Point.Title] = Point.Distance
	}

	infos := mysql.GetInfos("foodie", resultPlace)
	for idx, info := range infos {
		infos[idx].Distance = kvale[info.Title]

		uniq := strings.Split(info.Insta, "/")[4]
		infos[idx].Uniq = uniq
	}

	return c.JSON(http.StatusOK, infos)
}

// GetNearShopPlace func
func GetNearShopPlace(c echo.Context) error {
	subwayName := c.FormValue("subway")
	shopType := "shopping"

	subWayLatLng := mysql.GetSubwayLatLng(subwayName)
	if subWayLatLng.StationNm == "" {
		return c.String(http.StatusOK, "올바른 지하철역의 이름을 다시 입력해주세요")
	}
	placeList := mysql.GetPlaceLatLng(shopType)

	AllPointDis := utils.CalculateDistance(subwayName, subWayLatLng, placeList)

	var resultPlace []string
	kvale := map[string]float64{}

	for _, Point := range AllPointDis[:10] {
		resultPlace = append(resultPlace, Point.Title)
		kvale[Point.Title] = Point.Distance
	}

	infos := mysql.GetInfos("shopping", resultPlace)
	for idx, info := range infos {
		infos[idx].Distance = kvale[info.Title]

		uniq := strings.Split(info.Insta, "/")[4]
		infos[idx].Uniq = uniq
	}

	return c.JSON(http.StatusOK, infos)
}

// GetNearPlaceList func
func GetNearPlaceList(c echo.Context) error {
	subwayName := c.FormValue("subway")
	shopType := "placeList"

	subWayLatLng := mysql.GetSubwayLatLng(subwayName)
	if subWayLatLng.StationNm == "" {
		return c.String(http.StatusOK, "올바른 지하철역의 이름을 다시 입력해주세요")
	}
	placeList := mysql.GetPlaceLatLng(shopType)

	AllPointDis := utils.CalculateDistance(subwayName, subWayLatLng, placeList)

	var resultPlace []string
	kvale := map[string]float64{}

	for _, Point := range AllPointDis[:10] {
		resultPlace = append(resultPlace, Point.Title)
		kvale[Point.Title] = Point.Distance
	}

	infos := mysql.GetPlaceInfos(shopType, resultPlace)
	for idx, info := range infos {
		infos[idx].Distance = kvale[info.Title]

		uniq := strings.Split(info.Insta, "/")[4]
		infos[idx].Uniq = uniq
	}

	return c.JSON(http.StatusOK, infos)
}

// InsertFoodie func
func InsertFoodie(c echo.Context) error {
	place := new(m.Foodie)

	if err := c.Bind(place); err != nil {
		return c.String(http.StatusBadRequest, "request failed!")
	}

	res := mysql.InsertFoodie(place)
	return c.String(http.StatusOK, res)
}

// InsertShop func
func InsertShop(c echo.Context) error {
	place := new(m.Shopping)

	if err := c.Bind(place); err != nil {
		return c.String(http.StatusBadRequest, "request failed!")
	}

	res := mysql.InsertShop(place)
	return c.String(http.StatusOK, res)
}
