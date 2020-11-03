package handler

import (
	"fmt"
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

	fmt.Printf("Total Foodie's list count is %d \n", len(FoodieList))

	return c.JSON(http.StatusOK, FoodieList)
}

// GetShoppingList func
func GetShoppingList(c echo.Context) error {
	shoppingList := mysql.GetShoppingList()

	fmt.Printf("Total shopping's list count is %d \n", len(shoppingList))

	return c.JSON(http.StatusOK, shoppingList)
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
	placeList := mysql.GetPlaceLatLng(shopType)

	AllPointDis := utils.CalculateDistance(subwayName, subWayLatLng, placeList)

	var resultPlace []string
	kvale := map[string]float64{}

	for _, Point := range AllPointDis[:10] {
		resultPlace = append(resultPlace, Point.Title)
		kvale[Point.Title] = Point.Distance
	}

	infos := mysql.GetInfos("foodie",resultPlace)
	for idx, info := range infos {
		infos[idx].Distance = kvale[info.Title]
	}

	return c.JSON(http.StatusOK, infos)
}

// GetNearShopPlace func
func GetNearShopPlace(c echo.Context) error {
	subwayName := c.FormValue("subway")
	shopType := "shopping"

	subWayLatLng := mysql.GetSubwayLatLng(subwayName)
	placeList := mysql.GetPlaceLatLng(shopType)

	AllPointDis := utils.CalculateDistance(subwayName, subWayLatLng, placeList)

	var resultPlace []string
	kvale := map[string]float64{}

	for _, Point := range AllPointDis[:10] {
		resultPlace = append(resultPlace, Point.Title)
		kvale[Point.Title] = Point.Distance
	}

	infos := mysql.GetInfos("shopping",resultPlace)
	for idx, info := range infos {
		infos[idx].Distance = kvale[info.Title]
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
