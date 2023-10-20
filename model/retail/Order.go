package retail

import (
	"sync"
)

// Order 结构体
type Order struct {
	// 商品信息
	ItemList []ItemLineDto `json:"item_list,omitempty" xml:"item_list>item_line_dto,omitempty"`
	// 门店id
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 外部订单id
	OutOrderId string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// 提货类型
	ShippingType string `json:"shipping_type,omitempty" xml:"shipping_type,omitempty"`
	// 买家id
	BuyerId int64 `json:"buyer_id,omitempty" xml:"buyer_id,omitempty"`
	// 门店订单id
	StoreOrderId int64 `json:"store_order_id,omitempty" xml:"store_order_id,omitempty"`
	// 原价
	OriginPrice int64 `json:"origin_price,omitempty" xml:"origin_price,omitempty"`
	// 地址信息
	DeliveryAddress *DeliveryAddressDto `json:"delivery_address,omitempty" xml:"delivery_address,omitempty"`
	// 订单实付价格
	PayFee int64 `json:"pay_fee,omitempty" xml:"pay_fee,omitempty"`
}

var poolOrder = sync.Pool{
	New: func() any {
		return new(Order)
	},
}

// GetOrder() 从对象池中获取Order
func GetOrder() *Order {
	return poolOrder.Get().(*Order)
}

// ReleaseOrder 释放Order
func ReleaseOrder(v *Order) {
	v.ItemList = v.ItemList[:0]
	v.StoreId = ""
	v.OutOrderId = ""
	v.ShippingType = ""
	v.BuyerId = 0
	v.StoreOrderId = 0
	v.OriginPrice = 0
	v.DeliveryAddress = nil
	v.PayFee = 0
	poolOrder.Put(v)
}
