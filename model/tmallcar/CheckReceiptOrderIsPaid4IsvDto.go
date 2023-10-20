package tmallcar

// CheckReceiptOrderIsPaid4isvDto 结构体
type CheckReceiptOrderIsPaid4isvDto struct {
	// 无
	PaidOrderItems []OrderItem4isvDto `json:"paid_order_items,omitempty" xml:"paid_order_items>order_item4isv_dto,omitempty"`
	// 门店自定义编码
	OuterShopId string `json:"outer_shop_id,omitempty" xml:"outer_shop_id,omitempty"`
	// 工单id
	ReceiptId int64 `json:"receipt_id,omitempty" xml:"receipt_id,omitempty"`
}
