package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"../utils"
	"github.com/ilibs/gosql"
)

//AppConfig ..
type AppConfig struct {
	DataBase gosql.Config
}

//Config ..
var Config AppConfig

func init() {
	b, err := ioutil.ReadFile("config.json")
	utils.CheckErr(err)
	jsonStr := string(b)
	err = json.Unmarshal([]byte(jsonStr), &Config)
	fmt.Println(Config)
	utils.CheckErr(err)
}
