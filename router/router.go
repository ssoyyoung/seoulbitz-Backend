package router

import (
	"net/http"

	echo "github.com/labstack/echo"
	middleware "github.com/labstack/echo/middleware"
	handler "github.com/ssoyyoung.p/seoulbitz-Backend/handler"
)

// Router function
func Router() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())                             //Setting logger
	e.Use(middleware.Recover())                            //Recover from panics anywhere in the chain
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{ //CORS Middleware
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Success!")
	})

	getData := e.Group("/getData")
	{
		getData.GET("/foodie", handler.GetFoodieList)
		getData.GET("/shop", handler.GetShoppingList)
		getData.GET("/place", handler.GetPlaceList)
		getData.GET("/subway", handler.GetSubwayList)
	}

	getNearPlace := e.Group("/getNear")
	{
		getNearPlace.POST("/foodie", handler.GetNearFoodiePlace) //FormValue:subway
		getNearPlace.POST("/shop", handler.GetNearShopPlace)     //FormValue:subway
		getNearPlace.POST("/place", handler.GetNearPlaceList)    //FormValue:subway
	}

	insertData := e.Group("/insert")
	{
		insertData.POST("/foodie", handler.InsertFoodie)
		insertData.POST("/shop", handler.InsertShop)
	}

	// TODO : update place

	return e
}
