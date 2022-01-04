package main

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"strings"

	"pg_flame/pkg/config"
	"pg_flame/pkg/html"
	"pg_flame/pkg/plan"

	"github.com/jackc/pgx/v4"
	"github.com/labstack/echo"
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

	e := echo.New()

	e.GET("/info", func(c echo.Context) error {
		return c.JSON(200, map[string]string{
			"version": version,
		})
	})

	e.Static("/assets", "./assets")
	e.File("/", "./assets/index.html")

	e.POST("/", func(c echo.Context) error {
		q := fmt.Sprintf("explain (analyze, buffers, format json) %s", c.FormValue("query"))
		var exp string
		err := conn.QueryRow(context.Background(), q).Scan(&exp)
		if err != nil {
			c.String(500, err.Error())
			return nil
		}

		p, err := plan.New(strings.NewReader(exp))
		if err != nil {
			c.String(501, err.Error())
			return nil
		}

		out := new(bytes.Buffer)
		err = html.Generate(out, p)
		if err != nil {
			c.String(502, err.Error())
			return nil
		}

		return c.HTML(200, out.String())
	})

	e.Start(":5000")
}
