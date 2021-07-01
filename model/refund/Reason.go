package refund

// Reason 结构体
type Reason struct {
	//
	ReasonId int64 `json:"reason_id,omitempty" xml:"reason_id,omitempty"`
	//
	ReasonText string `json:"reason_text,omitempty" xml:"reason_text,omitempty"`
	//
	ReasonTips string `json:"reason_tips,omitempty" xml:"reason_tips,omitempty"`
}
