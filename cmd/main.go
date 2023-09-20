package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"

	routes "janio-client-backend/routes"
)

var db *sql.DB

func main() {
	routes.SetupRoutes(db)

	resp, err := http.Get("http://localhost:8080/client/orders")
	if err != nil {
		log.Fatalln(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	fmt.Printf(sb)
}
