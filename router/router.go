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

	e.GET("/getFoddie", handler.GetFoddieList)
	e.GET("/getShop", handler.GetShoppingList)
	e.GET("/getSubway", handler.GetSubwayList)

	e.POST("/nearSubway/:subway/:type", handler.GetNearSubway)

	return e
}
