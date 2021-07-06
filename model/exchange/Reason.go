package exchange

// Reason 结构体
type Reason struct {
	// 拒绝原因内容
	ReasonText string `json:"reason_text,omitempty" xml:"reason_text,omitempty"`
	// 拒绝原因ID
	ReasonId int64 `json:"reason_id,omitempty" xml:"reason_id,omitempty"`
}
