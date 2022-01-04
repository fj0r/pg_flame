package main

import (
	"bytes"
	"net/http"
	"strings"

	"pg_flame/pkg/config"
	"pg_flame/pkg/html"
	"pg_flame/pkg/plan"

	"github.com/gin-gonic/gin"
)

var (
	version  = "dev"
	pgConfig = config.Init()
)

func main() {
	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello",
		})
	})

    r.Static("/assets", "./assets")
	r.StaticFile("/", "./assets/index.html")

	r.POST("/", func(c *gin.Context) {
	    q := c.PostForm("query")
		p, err := plan.New(strings.NewReader(q))
		if err != nil {
		}

        out := new(bytes.Buffer)
		err = html.Generate(out, p)
		if err != nil {
		}
		c.Writer.WriteHeader(http.StatusOK)
		c.Writer.Write(out.Bytes())
	})

	r.Run(":5000")
}

