package main

import (
	"strings"
	"context"
	"os"
	"fmt"
	"net/http"
	"encoding/json"
    "io/ioutil"

	"pg_flame/pkg/config"
	"pg_flame/pkg/html"
	"pg_flame/pkg/plan"

	"github.com/jackc/pgx/v4"
)

var (
	version  = "dev"
	pgConfig = config.Init()
)

func infoHandler(w http.ResponseWriter, r *http.Request) {
    result := make(map[string]interface{})
    result["version"] = version

    w.Header().Set("Content-Type", "application/json")
    json, _ := json.Marshal(result)
    w.Write(json)
}


func main() {
    pgstr := pgConfig.URL()
    conn, err := pgx.Connect(context.Background(), pgstr)
    println(fmt.Sprintf("connect to postgres: %s", pgstr))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

    http.HandleFunc("/info", infoHandler)
    http.ListenAndServe(":5000", nil)
}

