package main

import (
	"database/sql"
	"janio-client-backend/api"
	"janio-client-backend/routes"
)

var db *sql.DB

func main() {
	api.PushToAdminInterval()
	routes.SetupRoutes(db)
}
