package server

import (
	"gopkg.in/gin-gonic/gin.v1"
	"path/filepath"
)

func staticProcess(r *gin.Engine) {
	r.Static("/static", staticPath)
    r.StaticFile("/favicon.ico", filepath.Join(staticPath, "favicon.ico"))
}