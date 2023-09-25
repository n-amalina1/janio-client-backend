package main

import (
	"database/sql"
	"fmt"
	"janio-client-backend/api"
	"janio-client-backend/routes"
)

var db *sql.DB

func main() {
	fmt.Printf("Successfully Added new orders")
	api.PushToAdminInterval()
	routes.SetupRoutes(db)
}
