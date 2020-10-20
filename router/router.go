package router

import (
	"fmt"
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

	e.GET("/getFoddie", func(c echo.Context) error {
		foddieList := mysql.GetFoddieList()
		fmt.Println(len(foddieList))
		return c.String(http.StatusOK, "Done")
	})

	return e
}
