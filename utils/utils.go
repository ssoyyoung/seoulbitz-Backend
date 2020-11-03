package utils

import (
	"fmt"
	"math"
	"sort"
	"strconv"

	m "github.com/ssoyyoung.p/seoulbitz-Backend/model"
)

// CheckErr func
func CheckErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

// CalculateLatAndLng func
// reference : https://www.geodatasource.com/developers/go
func CalculateLatAndLng(lat1, lng1, lat2, lng2 float64, unit ...string) float64 {
	const PI float64 = 3.141592653589793

	radlat1 := float64(PI * lat1 / 180)
	radlat2 := float64(PI * lat2 / 180)

	theta := float64(lng1 - lng2)
	radtheta := float64(PI * theta / 180)

	dist := math.Sin(radlat1)*math.Sin(radlat2) + math.Cos(radlat1)*math.Cos(radlat2)*math.Cos(radtheta)

	if dist > 1 {
		dist = 1
	}

	dist = math.Acos(dist)
	dist = dist * 180 / PI
	dist = dist * 60 * 1.1515

	if len(unit) > 0 {
		if unit[0] == "K" {
			dist = dist * 1.609344
		} else if unit[0] == "N" {
			dist = dist * 0.8684
		}
	}

	return dist
}

// StrToFloat64 func
func StrToFloat64(val string) float64 {
	f64, err := strconv.ParseFloat(val, 64)
	CheckErr(err)

	return f64
}

// CalculateDistance func
func CalculateDistance(subwayName string, subWayLatLng m.Subway, placeList []m.PlaceLatLng) []m.TwoPointDistance {
	PointDis := m.TwoPointDistance{}
	AllPointDis := []m.TwoPointDistance{}

	for _, place := range placeList {
		distance := CalculateLatAndLng(
			StrToFloat64(subWayLatLng.XpointWgs),
			StrToFloat64(subWayLatLng.YpointWgs),
			StrToFloat64(place.XpointWgs),
			StrToFloat64(place.YpointWgs),
			"K",
		)

		PointDis.Title = place.Title
		PointDis.Distance = distance

		AllPointDis = append(AllPointDis, PointDis)
	}

	AllPointDis = OrderedValue(AllPointDis)

	return AllPointDis
}

// OrderedValue func
func OrderedValue(AllPointDis []m.TwoPointDistance) []m.TwoPointDistance {

	orderVal := make(map[float64]string)
	returnVal := []m.TwoPointDistance{}
	value := m.TwoPointDistance{}

	for _, val := range AllPointDis {
		orderVal[val.Distance] = val.Title
	}

	keys := make([]float64, 0)
	for k := range orderVal {
		keys = append(keys, k)
	}

	sort.Float64s(keys)

	for _, k := range keys {
		value.Title = orderVal[k]
		value.Distance = k

		returnVal = append(returnVal, value)
	}

	return returnVal

}
