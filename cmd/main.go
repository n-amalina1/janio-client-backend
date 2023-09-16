package main

import (
	"database/sql"

	routes "janio-client-backend/routes"
)

var db *sql.DB

func main() {
	routes.SetupRoutes(db)
}
