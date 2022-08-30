package nrt

// NrtSmsDto 结构体
type NrtSmsDto struct {
	// 短信验证码
	SmsCode string `json:"sms_code,omitempty" xml:"sms_code,omitempty"`
}
