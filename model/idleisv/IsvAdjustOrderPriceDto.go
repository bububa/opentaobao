package idleisv

// IsvAdjustOrderPriceDto 结构体
type IsvAdjustOrderPriceDto struct {
	// 最新价格，单位分；必选，金额>0
	NewPriceFee int64 `json:"new_price_fee,omitempty" xml:"new_price_fee,omitempty"`
	// 最新邮费，单位分；必选，金额>=0
	NewTransportFee int64 `json:"new_transport_fee,omitempty" xml:"new_transport_fee,omitempty"`
	// 订单id
	BizOrderId int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
}
