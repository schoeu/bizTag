package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
	"path/filepath"
)

var (
    port = ":8910"
	theme = "default"
	staticPrefix = "./themes/" + theme
	staticPath = filepath.Join(staticPrefix, "static/")
)

func main() {
	RunServer("")
}


func RunServer(customPort string) {
	router := gin.Default()
    if customPort != "" {
        port = customPort
    }

	routers(router)
	staticPro(router)
	router.Run(port)
}

type Login struct {
    User     string `form:"username" json:"username" binding:"required"`
    Password string `form:"password" json:"password" binding:"required"`
}


func routers(r *gin.Engine) {
	r.LoadHTMLGlob(filepath.Join(staticPrefix, "views/*"))

	//apiRouter := r.Group("/api")
	r.POST("/api/login", func(c *gin.Context) {
		var form Login
        // This will infer what binder to use depending on the content-type header.
        if c.Bind(&form) == nil {
            if form.User == "test" && form.Password == "123" {
               c.HTML(http.StatusOK, "main.tmpl", gin.H{
			   		"title": "login",
			   })
            } else {
                c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
            }
        }
	})


	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "main.tmpl", gin.H{
			"title": "psfe",
		})
	})
}

func staticPro(r *gin.Engine) {
	r.Static("/static", staticPath)
    r.StaticFile("/favicon.ico", filepath.Join(staticPath, "favicon.ico"))
}