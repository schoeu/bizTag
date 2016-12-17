package server

import (
	"gopkg.in/gin-gonic/gin.v1"
	"path/filepath"

	"config"
	_ "github.com/go-sql-driver/mysql"
)

var (
	port         = ":8910"
	staticPrefix = ""
	staticPath   = ""
)

func RunServer() {
	cfg := config.GetConf()

	cfgPort := cfg.BizPort
	if cfgPort != "" {
		port = cfgPort
	}

	staticPrefix = "../themes/" + cfg.BizTheme
	staticPath = filepath.Join(staticPrefix, "static/")

	router := gin.Default()

	routers(router)
	staticProcess(router)
	router.Run(port)
}
