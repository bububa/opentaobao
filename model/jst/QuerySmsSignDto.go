package jst

// QuerySmsSignDto 结构体
type QuerySmsSignDto struct {
	// 被拒绝的原因
	Reason string `json:"reason,omitempty" xml:"reason,omitempty"`
	// 申请的签名
	SignName string `json:"sign_name,omitempty" xml:"sign_name,omitempty"`
	// 签名创建时间
	CreateDate string `json:"create_date,omitempty" xml:"create_date,omitempty"`
	// 拒绝
	SignStatus int64 `json:"sign_status,omitempty" xml:"sign_status,omitempty"`
}
