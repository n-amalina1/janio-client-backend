package models

var item = ClientItem{
	ItemDescription: "Chanel 22 Bag",
	ItemCategory:    "Fashin",
	ItemSku:         "ITEMBAG22",
	ItemQuantity:    3,
	ItemPrice:       12000,
	ItemCurrency:    "SGD",
}

var order1 = ClientOrder{
	OrderDetails: OrderDetails{
		OrderLength: 20,
		OrderWidth:  20,
		OrderHeight: 20,
		OrderWeight: 20,
	},
	Address: Address{
		Consignee: ClientConsignee{
			ConsigneeName:        "Jennie",
			ConsigneePhoneNumber: "+62068922589",
			ConsigneeCountry:     "Indonesia",
			ConsigneeAddress:     "Jl RA Kartini 59 20635",
			ConsigneePostal:      20635,
			ConsigneeState:       "Medan",
			ConsigneeCity:        "Tebingtinggi Deli",
			ConsigneeProvince:    "Medan",
			ConsigneeEmail:       "jenniekim@yahoo.com",
		},
		Pickup: ClientPickup{
			PickupName:        "Stella",
			PickupPhoneNumber: "+65 68361577",
			PickupCountry:     "Singapore",
			PickupAddress:     "Hotel Grand Central 22 Cavenagh Road #01-01",
			PickupPostal:      229617,
			PickupState:       "Singapore",
			PickupCity:        "Singapore",
			PickupProvince:    "null",
		},
	},
	Items: []ClientItem{item},
}

var order2 = ClientToDBOrder{

	OrderLength: 20,
	OrderWidth:  20,
	OrderHeight: 20,
	OrderWeight: 20,
	OrderStatus: "Pending",
	Consignee: ClientConsignee{
		ConsigneeName:        "Sally",
		ConsigneePhoneNumber: "+600320937144",
		ConsigneeCountry:     "Malaysia",
		ConsigneeAddress:     "No. 20 Mdn Setia 2 Bukit Damansara",
		ConsigneePostal:      50490,
		ConsigneeState:       "Wilayah Persekutuan",
		ConsigneeCity:        "Kuala Lumpur",
		ConsigneeProvince:    "Wilayah Persekutuan",
		ConsigneeEmail:       "salllyyy@gmail.com",
	},
	Pickup: ClientPickup{
		PickupName:        "Stella",
		PickupPhoneNumber: "+6564789462",
		PickupCountry:     "Singapore",
		PickupAddress:     "630 Hougang Ave 8 #01-K1",
		PickupPostal:      530630,
		PickupState:       "null",
		PickupCity:        "Singapore",
		PickupProvince:    "null",
	},
	Items: []ClientItem{item},
}

var OrdersInit = []ClientOrder{order1}
var OrdersDB = []ClientToDBOrder{order2}
