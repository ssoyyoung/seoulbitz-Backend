package svc

import (
	"fmt"

	m "github.com/ssoyyoung.p/seoulbitz-Backend/model"
)

// DistanceCalculate func
func DistanceCalculate(subWayLatLng m.Subway, placeList []m.PlaceLatLng) {
	fmt.Println(placeList)
	fmt.Println(subWayLatLng)
}
