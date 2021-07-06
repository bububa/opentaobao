package wdk

// Reasonlist 结构体
type Reasonlist struct {
	// 退款原因说明
	ReasonText string `json:"reason_text,omitempty" xml:"reason_text,omitempty"`
	// 退款原因ID
	ReasonId int64 `json:"reason_id,omitempty" xml:"reason_id,omitempty"`
}
