package routes

import (
	"database/sql"
	"net/http"

	models "janio-client-backend/models"

	"github.com/gin-gonic/gin"
)

var db *sql.DB

func SetupRoutes(d *sql.DB) {
	db = d
	router := gin.Default()
	router.GET("new/orders", GetOrdersClient)

	router.Run("localhost:8008")
}

func GetOrdersClient(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.OrdersInit)
}
