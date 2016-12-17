package server

import (
	"database/sql"
	"fmt"
	"gopkg.in/gin-gonic/gin.v1"
	"log"
	"net/http"
	"net/url"
	"path/filepath"
)

type Login struct {
	User     string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type Signup struct {
	User     string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required"`
}

type AddSite struct {
	SiteName  string `form:"site_name" json:"site_name" binding:"required"`
	SiteUrl   string `form:"site_url" json:"site_url" binding:"required"`
	SiteGroup string `form:"site_group" json:"email"`
}

func routers(r *gin.Engine) {

	r.LoadHTMLGlob(filepath.Join(staticPrefix, "views/*"))

	db := getDB()

	// 主页
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "main.tmpl", gin.H{
			"title":    "psfe",
			"username": "schoeu",
		})
	})

	// 注册GET
	r.GET("/signup", func(c *gin.Context) {
		c.HTML(http.StatusOK, "signup.tmpl", gin.H{
			"title":   "Sign up",
			"isLogin": false,
		})
	})

	// 注册POST
	r.POST("/signup", func(c *gin.Context) {
		var form Signup
		if c.Bind(&form) == nil {
			var id string
			uname := form.User
			rows, err := db.Query("select id from users where username = ?", uname)
			defer rows.Close()

			for rows.Next() {
				err := rows.Scan(&id)
				checkErr(err)
			}

			err = rows.Err()
			checkErr(err)

			// 表中无记录
			if id == "" {
				stmt, err := db.Prepare("insert into users(username, password, email)values(?,?,?)")
				checkErr(err)

				defer stmt.Close()

				_, err = stmt.Exec(uname, form.Password, form.Email)
				if err != sql.ErrNoRows {
					c.JSON(http.StatusOK, gin.H{
						"errorNo":  0,
						"has":      0,
						"username": uname,
					})
				} else {
					checkErr(err)
				}
			} else {
				checkErr(err)
				c.JSON(http.StatusOK, gin.H{
					"errorNo": 0,
					"has":     1,
				})
			}

			checkErr(err)
		}
	})

	// 登录GET
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.tmpl", gin.H{
			"title":   "Sign in",
			"isLogin": false,
		})
	})

	// 登录POST
	r.POST("/login", func(c *gin.Context) {
		var form Login
		if c.Bind(&form) == nil {
			var psw string
			rows := db.QueryRow("select password from users where username = ?", form.User)

			err := rows.Scan(&psw)
			if err == sql.ErrNoRows {
				c.JSON(http.StatusOK, gin.H{
					"errorNo": 0,
					"issigup": 0,
				})
				return
			}

			checkErr(err)

			if form.Password == psw {
				c.Redirect(http.StatusFound, "/")
			} else {
				c.JSON(http.StatusOK, gin.H{
					"errorNo": 0,
					"issigup": 1,
					"msg":     "wrong password.",
				})
			}
		}
	})

	// 注册POST
	r.POST("/addsite", func(c *gin.Context) {
		var form AddSite
		if c.Bind(&form) == nil {

			var id string
			siteUrl := form.SiteUrl
			siteInfo, err := url.Parse(siteUrl)
			checkErr(err)
			fmt.Println(siteInfo)

			host := siteInfo.Host

			siteFullUrl := siteInfo.Scheme + "://" + siteInfo.Host

			siteIcon := siteUrl + "/favicon.ico"

			rows, err := db.Query("select id from sites where site_name = ?", host)
			defer rows.Close()

			for rows.Next() {
				err := rows.Scan(&id)
				checkErr(err)
			}

			err = rows.Err()
			checkErr(err)

			// 表中无记录
			if id == "" {
				stmt, err := db.Prepare("insert into sites(site_url, site_name, site_group, site_icon)values(?,?,?,?)")
				checkErr(err)

				defer stmt.Close()

				_, err = stmt.Exec(siteFullUrl, form.SiteName, form.SiteGroup, siteIcon)
				if err != sql.ErrNoRows {
					c.JSON(http.StatusOK, gin.H{
						"errorNo": 0,
						"ok":      1,
					})
				} else {
					checkErr(err)
				}
			} else {
				checkErr(err)
				c.JSON(http.StatusOK, gin.H{
					"errorNo": 0,
					"has":     1,
				})
			}

			checkErr(err)
		}
	})
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}

}
