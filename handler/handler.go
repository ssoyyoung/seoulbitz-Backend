package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	m "github.com/ssoyyoung.p/seoulbitz-Backend/model"
	mysql "github.com/ssoyyoung.p/seoulbitz-Backend/mysql"
	utils "github.com/ssoyyoung.p/seoulbitz-Backend/utils"
)

// GetFoddieList func
func GetFoddieList(c echo.Context) error {
	foddieList := mysql.GetFoddieList()

	fmt.Printf("Total foddie's list count is %d \n", len(foddieList))

	return c.JSON(http.StatusOK, foddieList)
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

// GetNearSubway func
func GetNearSubway(c echo.Context) error {
	subwayName := c.Param("subway")
	shopType := c.Param("type")

	subWayLatLng := mysql.GetSubwayLatLng(subwayName)
	placeList := mysql.GetPlaceLatLng(shopType)

	PointDis := m.TwoPointDistance{}
	AllPointDis := []m.TwoPointDistance{}

	for _, place := range placeList {
		distance := utils.CalculateLatAndLng(
			utils.StrToFloat64(subWayLatLng.XpointWgs),
			utils.StrToFloat64(subWayLatLng.YpointWgs),
			utils.StrToFloat64(place.XpointWgs),
			utils.StrToFloat64(place.YpointWgs),
			"K",
		)

		PointDis.Subway = subwayName
		PointDis.Destination = place.Title
		PointDis.Distance = distance

		AllPointDis = append(AllPointDis, PointDis)
	}

	return c.JSON(http.StatusOK, AllPointDis)
}

// InsertFoddie func
func InsertFoddie(c echo.Context) error {
	place := new(m.Foddie)

	if err := c.Bind(place); err != nil {
		return c.String(http.StatusBadRequest, "request failed!")
	}

	res := mysql.InsertFoddie(place)
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
