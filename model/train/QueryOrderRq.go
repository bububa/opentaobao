package train

// QueryOrderRq 结构体
type QueryOrderRq struct {
	// 主单号
	TtpOrderId int64 `json:"ttp_order_id,omitempty" xml:"ttp_order_id,omitempty"`
}
