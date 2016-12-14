package server

import (
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
	"path/filepath"
	"fmt"
)

func routers(r *gin.Engine) {
	r.LoadHTMLGlob(filepath.Join(staticPrefix, "views/*"))

	db := getDB()
    defer db.Close()

	// Execute the query
    rows, err := db.Query("SELECT * FROM vs_users")
    if err != nil {
        panic(err.Error()) 
    }

	// Get column names
    columns, err := rows.Columns()
    if err != nil {
        panic(err.Error()) 
    }

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