package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"janio-client-backend/models"
	"log"
	"net/http"
	"time"
)

func FormatClientToDB(orders []models.ClientOrder) []models.ClientToDBOrder {
	var ordersC []models.ClientToDBOrder

	for _, o := range orders {
		var order models.ClientToDBOrder

		order.OrderLength = o.OrderDetails.OrderLength
		order.OrderWidth = o.OrderDetails.OrderWidth
		order.OrderHeight = o.OrderDetails.OrderHeight
		order.OrderWeight = o.OrderDetails.OrderWeight
		order.OrderStatus = "Pending"
		order.Consignee = o.Address.Consignee
		order.Pickup = o.Address.Pickup
		order.Items = o.Items

		ordersC = append(ordersC, order)
	}
	return ordersC
}

func pushToAdmin() {
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
	fmt.Printf("%s\n", sb)
	fmt.Printf("Successfully Added new orders")
}

func PushToAdminInterval() {
	ticker := time.NewTicker(1 * time.Hour)

	done := make(chan bool)
	go func() {
		for {
			select {
			case <-ticker.C:

				pushToAdmin()
			case <-done:

				ticker.Stop()
				return
			}
		}
	}()

	time.Sleep(10 * time.Hour)
	done <- true
}
