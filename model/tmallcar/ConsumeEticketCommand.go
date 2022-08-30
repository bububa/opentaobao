package tmallcar

// ConsumeEticketCommand 结构体
type ConsumeEticketCommand struct {
	// 核销码
	EticketCode string `json:"eticket_code,omitempty" xml:"eticket_code,omitempty"`
	// 订单id
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 门店id
	StoreId int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
}
