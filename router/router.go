package router

import (
	"net/http"

	echo "github.com/labstack/echo"
	handler "github.com/ssoyyoung.p/seoulbitz-Backend/handler"
)

// Router function
func Router() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Success!")
	})

	getData := e.Group("/getData")
	{
		getData.GET("/foodie", handler.GetFoodieList)
		getData.GET("/shop", handler.GetShoppingList)
		getData.GET("/subway", handler.GetSubwayList)
	}

	getNearPlace := e.Group("/getNear")
	{
		getNearPlace.POST("/foodie/:subway", handler.GetNearFoodiePlace)
		getNearPlace.POST("/shop/:subway", handler.GetNearShopPlace)
	}

	insertData := e.Group("/insert")
	{
		insertData.POST("/foodie", handler.InsertFoodie)
		insertData.POST("/shop", handler.InsertShop)
	}

	return e
}
