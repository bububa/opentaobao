package btrip

// ReturnReasonDetail 结构体
type ReturnReasonDetail struct {
	// 原因的文案展示
	ReasonShow string `json:"reason_show,omitempty" xml:"reason_show,omitempty"`
	// 退票原因注释
	ExtendDesc string `json:"extend_desc,omitempty" xml:"extend_desc,omitempty"`
	// 原因的类型
	ReasonType int64 `json:"reason_type,omitempty" xml:"reason_type,omitempty"`
	// 原因code
	ReasonCode int64 `json:"reason_code,omitempty" xml:"reason_code,omitempty"`
	// 自愿或非自愿 1-非自愿、0-自愿
	Volunteer int64 `json:"volunteer,omitempty" xml:"volunteer,omitempty"`
	// 个人原因或航司原因  0-航司、1-个人
	Person int64 `json:"person,omitempty" xml:"person,omitempty"`
}
