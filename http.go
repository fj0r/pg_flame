package main

import (
	"bytes"
	"net/http"
	"strings"
	"context"
	"os"
	"fmt"

	"pg_flame/pkg/config"
	"pg_flame/pkg/html"
	"pg_flame/pkg/plan"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

var (
	version  = "dev"
	pgConfig = config.Init()
)

func main() {
    pgstr := pgConfig.URL()
    conn, err := pgx.Connect(context.Background(), pgstr)
    println(fmt.Sprintf("connect to postgres: %s", pgstr))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	r := gin.Default()

	r.GET("/info", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"version": version,
		})
	})

    r.Static("/assets", "./assets")
	r.StaticFile("/", "./assets/index.html")

	r.POST("/", func(c *gin.Context) {
	    q := fmt.Sprintf("explain (analyze, buffers, format json) %s", c.PostForm("query"))
	    var exp string
	    err := conn.QueryRow(context.Background(), q).Scan(&exp)
	    if err != nil {
	        c.String(500, err.Error())
	        return
	    }


		p, err := plan.New(strings.NewReader(exp))
		if err != nil {
	        c.String(501, err.Error())
	        return
		}

        out := new(bytes.Buffer)
		err = html.Generate(out, p)
		if err != nil {
	        c.String(502, err.Error())
	        return
		}

		c.Writer.WriteHeader(http.StatusOK)
		c.Writer.Write(out.Bytes())
	})

	r.Run(":5000")
}

