package logistic

// OrderDto 结构体
type OrderDto struct {
	// products
	Products []ProductDto `json:"products,omitempty" xml:"products>product_dto,omitempty"`
	// Shipment order id created. AE will save relationship with logistics provider&#39;s shipment order_id
	LogisticsOrderId string `json:"logistics_order_id,omitempty" xml:"logistics_order_id,omitempty"`
	// Logistics provider Id of the shipping order such as 1-Cainiao, 2-Pegaki, 3-Frenet, 4-Delivery Hub, etc.
	LogisticsChannelOrderId string `json:"logistics_channel_order_id,omitempty" xml:"logistics_channel_order_id,omitempty"`
	// tracking code
	TrackingCode string `json:"tracking_code,omitempty" xml:"tracking_code,omitempty"`
	// Delivery option id selected by a sender
	ShippingMethodId string `json:"shipping_method_id,omitempty" xml:"shipping_method_id,omitempty"`
	// Logistics provider Id of the shipping order such as 1-Cainiao, 2-Pegaki, 3-Frenet, 4-Delivery Hub, etc.
	LogisticsChannelId string `json:"logistics_channel_id,omitempty" xml:"logistics_channel_id,omitempty"`
	// status of logistics order,1-NEW 2-CREATED 3-SHIPPED 4-IN TRANSIT 5-DELIVERY FAILURE 6-DELIVERED 7-CANCELLED 8-UNKNOWN
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// tracking url for senders to check status
	TrackingUrl string `json:"tracking_url,omitempty" xml:"tracking_url,omitempty"`
	// logistics channel name
	LogisticsChannelName string `json:"logistics_channel_name,omitempty" xml:"logistics_channel_name,omitempty"`
	// parcel
	Parcel *ParcelDto `json:"parcel,omitempty" xml:"parcel,omitempty"`
	// It&#39;s sales order id and a 16-digit number. There will be mapping between trade order and shipment order
	TradeOrderId int64 `json:"trade_order_id,omitempty" xml:"trade_order_id,omitempty"`
	// receiver
	ReceiptAddress *AddressDto `json:"receipt_address,omitempty" xml:"receipt_address,omitempty"`
	// sender
	SenderAddress *AddressDto `json:"sender_address,omitempty" xml:"sender_address,omitempty"`
	// invoice
	Invoice *InvoiceDto `json:"invoice,omitempty" xml:"invoice,omitempty"`
}
