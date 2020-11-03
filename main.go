package main

import (
	"fmt"

	router "github.com/ssoyyoung.p/seoulbitz-Backend/router"
)

func main() {
	debug := false
	echoR := router.Router()

	fmt.Println("Start echo server..")

	if debug {
		echoR.Logger.Fatal(echoR.Start(":1323"))
	} else {
		echoR.Logger.Fatal(echoR.StartTLS(":1323", "cert.pem", "privkey.pem"))
	}
}
