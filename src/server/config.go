package server

import (
	"encoding/json"
	"io/ioutil"
	"fmt"
)

var (
	defaultPath   = "../config/config.json"
)

type cfg struct {
	DBAddress string `json:"db_address"`
    DBName string `json:"db_name"`
    DBPassword string `json:"db_password"`
    DBUsername string `json:"db_username"`
}


// 获取配置文件
func getConf() cfg{
    var c cfg
    content, err := ioutil.ReadFile(defaultPath)
    fmt.Println(string(content))
    if err != nil {
        panic(err)
    }
    json.Unmarshal(content, &c)
	return c
}
