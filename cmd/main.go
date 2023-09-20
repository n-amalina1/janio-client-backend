package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"janio-client-backend/models"
	routes "janio-client-backend/routes"
)

var db *sql.DB

func main() {
	routes.SetupRoutes(db)
	someFunction()

}

func someFunction() {
	bodyBytes, err := json.Marshal(models.OrdersDB)
	if err != nil {
		log.Fatal(err)
	}
	reader := bytes.NewReader(bodyBytes)

	resp, err := http.Post("http://localhost:8080/client/orders", "application/json", reader)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	fmt.Printf(sb)
}
