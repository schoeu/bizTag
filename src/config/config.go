package config

import (
	"encoding/json"
	"io/ioutil"
)

var (
	defaultPath = "../config/config.json"
)

type cfg struct {
	DBAddress  string `json:"db_address"`
	DBName     string `json:"db_name"`
	DBPassword string `json:"db_password"`
	DBUsername string `json:"db_username"`
	QiniuAK string `json:"qiniu_ak"`
	QiniuSK string `json:"qiniu_sk"`
	QiniuBucket string `json:"qiniu_bucket"`
}

var c cfg

// 获取配置文件
func ReadConf() cfg {
	content, err := ioutil.ReadFile(defaultPath)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(content, &c)
}

func GetConf() cfg{
	return c
}