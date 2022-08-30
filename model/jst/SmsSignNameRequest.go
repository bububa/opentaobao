package jst

// SmsSignNameRequest 结构体
type SmsSignNameRequest struct {
	// 短信签名
	SignName string `json:"sign_name,omitempty" xml:"sign_name,omitempty"`
	// 描述信息
	Description string `json:"description,omitempty" xml:"description,omitempty"`
}
