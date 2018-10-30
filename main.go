package main

import (
	"fmt"
	"net/http"
	"test/conf"

	_ "./routers"
	"./utils"
	"github.com/ilibs/gosql"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err, "========")
		}
	}()

	//数据库初始化
	configs := make(map[string]*gosql.Config)
	configs["default"] = &conf.Config.DataBase
	gosql.Connect(configs)

	err := http.ListenAndServe(":3333", nil)
	utils.CheckErr(err)
}
