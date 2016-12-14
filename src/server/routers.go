package server

import (
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
	"path/filepath"
)

func routers(r *gin.Engine) {
	r.LoadHTMLGlob(filepath.Join(staticPrefix, "views/*"))

	db := getDB()
    defer db.Close()

	r.POST("/api/login", func(c *gin.Context) {
		var form Login
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