package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
)

var (
    port = ":8910"
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

	router.Run(port)
}

func routers(r *gin.Engine) {
	r.GET("/a/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello %s", "test")
	})
}

func staticPro(r *gin.Engine) {

}