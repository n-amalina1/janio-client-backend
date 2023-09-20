package api

import (
	"janio-client-backend/models"
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
