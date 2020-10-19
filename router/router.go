package router

import (
	"net/http"

	echo "github.com/labstack/echo"
	mysql "github.com/ssoyyoung.p/seoulbitz-Backend/mysql"
)

// Router function
func Router() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Success!")
	})

	e.GET("/mysql", func(c echo.Context) error {
		mysql.ConnectDB()
		return c.String(http.StatusOK, "DB connection test!")
	})

	return e
}
