package utils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

//ToJSON ..
func ToJSON(obj interface{}) string {
	data, err := json.Marshal(obj)
	CheckErr(err)
	json := string(data[:])

	return json
}

//CheckErr ..
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

//Encrypt ..
func Encrypt(p string) string {
	h := md5.New()
	h.Write([]byte(p))
	return hex.EncodeToString(h.Sum(nil))
}

//InterfaceToStr ..
func InterfaceToStr(obj interface{}) string {
	msg := fmt.Sprint(obj)
	return msg
}
