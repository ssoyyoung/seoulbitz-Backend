package main

import (
	router "github.com/ssoyyoung.p/seoulbitz-Backend/router"
)

func main() {
	echoR := router.Router()

	echoR.Logger.Fatal(echoR.Start(":1323"))
}
