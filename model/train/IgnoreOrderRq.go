package train

// IgnoreOrderRq 结构体
type IgnoreOrderRq struct {
	// 子单号 （不填默认全部忽略）
	SubOrderId []int64 `json:"sub_order_id,omitempty" xml:"sub_order_id>int64,omitempty"`
	// 主单号
	TtpOrderId int64 `json:"ttp_order_id,omitempty" xml:"ttp_order_id,omitempty"`
}
