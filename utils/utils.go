package utils

import "log"

// CheckErr func
func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
