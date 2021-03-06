package retail

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
