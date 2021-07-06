package maitix

// QueryPayResultParam 结构体
type QueryPayResultParam struct {
	// 订单金额
	Amount string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 订单日期
	Date string `json:"date,omitempty" xml:"date,omitempty"`
	// 订单号
	OrderNo string `json:"order_no,omitempty" xml:"order_no,omitempty"`
	// 大麦订单号
	DamaiOrderNo int64 `json:"damai_order_no,omitempty" xml:"damai_order_no,omitempty"`
}
