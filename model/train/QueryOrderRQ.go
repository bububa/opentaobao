package train

// QueryOrderRQ 结构体
type QueryOrderRQ struct {
	// 主单号
	TtpOrderId int64 `json:"ttp_order_id,omitempty" xml:"ttp_order_id,omitempty"`
}
