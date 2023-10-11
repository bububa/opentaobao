package logistic

// TmsSmsInfoRequest 结构体
type TmsSmsInfoRequest struct {
	// 短信发送时间。smsInfo有值时，必需
	SmsSendTime string `json:"sms_send_time,omitempty" xml:"sms_send_time,omitempty"`
	// 短信信息，smsInfo有值时，且smsSendStatus = true时，必需
	SmsText string `json:"sms_text,omitempty" xml:"sms_text,omitempty"`
	// 短信状态。true，发送成功；false，发送失败
	SmsSendStatus bool `json:"sms_send_status,omitempty" xml:"sms_send_status,omitempty"`
}
