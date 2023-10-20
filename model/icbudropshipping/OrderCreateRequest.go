package icbudropshipping

import (
	"sync"
)

// OrderCreateRequest 结构体
type OrderCreateRequest struct {
	// Product list
	ProductList []TradeEcologyOrderCreateProduct `json:"product_list,omitempty" xml:"product_list>trade_ecology_order_create_product,omitempty"`
	// Provide the order number corresponding to the 3rd party ISV
	ChannelReferId string `json:"channel_refer_id,omitempty" xml:"channel_refer_id,omitempty"`
	// Put the order number provided by the 3rd party platform and the name of the 3rd party platform. For example, if the order number is for a transaction made on Shopify, put “Shopify” and the order number. &lt;br /&gt;  Platform Names can be case ignored:&lt;br /&gt; Shopify,CommerceHQ,WooCommerce,GrooveKart,BigCommerce
	Properties string `json:"properties,omitempty" xml:"properties,omitempty"`
	// order remark
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// Logistics details
	LogisticsDetail *LogisticsDetail `json:"logistics_detail,omitempty" xml:"logistics_detail,omitempty"`
	// Payment details
	PaymentDetail *PaymentDetail `json:"payment_detail,omitempty" xml:"payment_detail,omitempty"`
}

var poolOrderCreateRequest = sync.Pool{
	New: func() any {
		return new(OrderCreateRequest)
	},
}

// GetOrderCreateRequest() 从对象池中获取OrderCreateRequest
func GetOrderCreateRequest() *OrderCreateRequest {
	return poolOrderCreateRequest.Get().(*OrderCreateRequest)
}

// ReleaseOrderCreateRequest 释放OrderCreateRequest
func ReleaseOrderCreateRequest(v *OrderCreateRequest) {
	v.ProductList = v.ProductList[:0]
	v.ChannelReferId = ""
	v.Properties = ""
	v.Remark = ""
	v.LogisticsDetail = nil
	v.PaymentDetail = nil
	poolOrderCreateRequest.Put(v)
}
