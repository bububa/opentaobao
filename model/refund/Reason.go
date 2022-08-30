package refund

// Reason 结构体
type Reason struct {
	// 退款原因文案
	ReasonText string `json:"reason_text,omitempty" xml:"reason_text,omitempty"`
	// 已经影响商品完好
	ReasonTips string `json:"reason_tips,omitempty" xml:"reason_tips,omitempty"`
	// 退款原因ID
	ReasonId int64 `json:"reason_id,omitempty" xml:"reason_id,omitempty"`
}
