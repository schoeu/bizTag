package server

import (
	"gopkg.in/gin-gonic/gin.v1"
	"log"
	"net/http"
	"path/filepath"
	"database/sql"
)

type Login struct {
	User     string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type Signup struct {
	User     string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Email string `form:"email" json:"email" binding:"required"`
}



func routers(r *gin.Engine) {

	r.LoadHTMLGlob(filepath.Join(staticPrefix, "views/*"))

	db := getDB()

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
		var form Signup
		if c.Bind(&form) == nil {
			var id string
			rows, err := db.Query("select id from users where username = ?", form.User)
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

				_, err = stmt.Exec(form.User, form.Password, form.Email)
				if err != sql.ErrNoRows {
					c.JSON(http.StatusOK, gin.H{
						"errorNo": 0,
						"has": 0,
					})
				} else {
					checkErr(err)
				}
			} else {
				checkErr(err)
				c.JSON(http.StatusOK, gin.H{
					"errorNo": 0,
					"has": 1,
				})
			}

			checkErr(err)
		}
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
					"msg": "wrong password.",
				})
			}
		}
	})
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}

}
