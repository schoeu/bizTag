package server

import (
	"gopkg.in/gin-gonic/gin.v1"
	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
)

var (
    port = ":8911"
	theme = "default"
	staticPrefix = "../themes/" + theme
	staticPath = filepath.Join(staticPrefix, "static/")
)

func RunServer(customPort string) {
	router := gin.Default()
    if customPort != "" {
        port = customPort
    }	

	routers(router)
	staticProcess(router)
	router.Run(port)
}

type Login struct {
    User     string `form:"username" json:"username" binding:"required"`
    Password string `form:"password" json:"password" binding:"required"`
}

