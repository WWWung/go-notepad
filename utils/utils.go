package utils

import (
	"encoding/json"
	"test/throw"
)

//ToJSON ..
func ToJSON(obj interface{}) string {
	data, err := json.Marshal(obj)
	throw.CheckErr(err)
	json := string(data[:])

	return json
}

//CheckErr ..
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
