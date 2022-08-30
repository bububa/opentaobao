package logistic

// CreateOrderRequestTopDto 结构体
type CreateOrderRequestTopDto struct {
	// products info
	Products []ProductTopDto `json:"products,omitempty" xml:"products>product_top_dto,omitempty"`
	// Logistics provider Id of the shipping order such as 1-Cainiao, 2-Pegaki, 3-Frenet, 4-Delivery Hub, etc.
	LogisticsChannelId string `json:"logistics_channel_id,omitempty" xml:"logistics_channel_id,omitempty"`
	// Input delivery_cost in logistics method query API. If no value returned, then input 1
	ProviderShippingCosts string `json:"provider_shipping_costs,omitempty" xml:"provider_shipping_costs,omitempty"`
	// Find it in responses of shipping method query API
	QueryId string `json:"query_id,omitempty" xml:"query_id,omitempty"`
	// sender address
	SenderAddress *AddressTopDto `json:"sender_address,omitempty" xml:"sender_address,omitempty"`
	// parcel info
	Parcel *ParcelTopDto `json:"parcel,omitempty" xml:"parcel,omitempty"`
	// Delivery option id selected by a sender
	DeliveryMethodId int64 `json:"delivery_method_id,omitempty" xml:"delivery_method_id,omitempty"`
	// It's sales order id and a 16-digit number. There will be mapping between trade order and shipment order
	TradeOrderId int64 `json:"trade_order_id,omitempty" xml:"trade_order_id,omitempty"`
	// receiver address
	ReceiptAddress *AddressTopDto `json:"receipt_address,omitempty" xml:"receipt_address,omitempty"`
	// invoice info
	Invoice *InvoiceTopDto `json:"invoice,omitempty" xml:"invoice,omitempty"`
}
