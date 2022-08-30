package tmallgenie

// MemberChargeRequest 结构体
type MemberChargeRequest struct {
	// 加密的手机号，采用RSA加密，由猫精提供加密方法
	EncryptedMobile string `json:"encrypted_mobile,omitempty" xml:"encrypted_mobile,omitempty"`
	// 合作方自己生成需要保证唯一，至少16位。
	OrderNo string `json:"order_no,omitempty" xml:"order_no,omitempty"`
	// 猫精提供的bizCode
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// 加密后的签名，签名方法见：猫精提供文档
	Sign string `json:"sign,omitempty" xml:"sign,omitempty"`
	// 当前请求的时间戳，例如：1641468394035
	Timestamp int64 `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
}
