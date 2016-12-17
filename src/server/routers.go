package server

import (
	"gopkg.in/gin-gonic/gin.v1"
	"log"
	"net/http"
	"path/filepath"
)

type Login struct {
	User     string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func routers(r *gin.Engine) {
	r.LoadHTMLGlob(filepath.Join(staticPrefix, "views/*"))

	// 主页
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "main.tmpl", gin.H{
			"title": "psfe",
		})
	})

	// 注册GET
	r.GET("/signup", func(c *gin.Context) {
		c.HTML(http.StatusOK, "signup.tmpl", gin.H{
			"title": "Sign up",
		})
	})

	// 注册POST
	r.POST("/signup", func(c *gin.Context) {

	})

	// 登录GET
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.tmpl", gin.H{
			"title": "Sign in",
		})
	})

	// 登录POST
	r.POST("/login", func(c *gin.Context) {
		var form Login
		db := getDB()
		defer db.Close()
		if c.Bind(&form) == nil {

			err := db.Ping()
			if err != nil {
				log.Fatal(err)
			}

			rows, err := db.Query("select password from users where username = ?", form.User)
			if err != nil {
				log.Println(err)
			}

			defer rows.Close()
			var psw string
			for rows.Next() {
				err := rows.Scan(&psw)
				if err != nil {
					log.Fatal(err)
				}
			}

			err = rows.Err()
			if err != nil {
				log.Fatal(err)
			}

			if form.Password == psw {
				c.Redirect(http.StatusFound, "/")
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			}
		}
	})
}
