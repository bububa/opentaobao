package logistic

// QueryShippingMethodRequestTopDto 结构体
type QueryShippingMethodRequestTopDto struct {
	// product list
	Products []ProductTopDto `json:"products,omitempty" xml:"products>product_top_dto,omitempty"`
	// It&#39;s sales order id and a 16-digit number To confirm logistics service provider in the order such as Cainiao, Pegaki, Frenet, Delivery Hub, etc.
	TradeOrderId int64 `json:"trade_order_id,omitempty" xml:"trade_order_id,omitempty"`
	// origin address
	SenderAddress *AddressTopDto `json:"sender_address,omitempty" xml:"sender_address,omitempty"`
	// extension
	Extension *ExtensionTopDto `json:"extension,omitempty" xml:"extension,omitempty"`
	// destination address
	ReceiptAddress *AddressTopDto `json:"receipt_address,omitempty" xml:"receipt_address,omitempty"`
}
