package alitripmerchant

// MemberCardParamVO 结构体
type MemberCardParamVO struct {
	// 微信公众号会员卡模板 id
	CardId string `json:"card_id,omitempty" xml:"card_id,omitempty"`
	// 用户会员卡 卡号
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 当前用户在微信公众号下的OpenId
	OpenId string `json:"open_id,omitempty" xml:"open_id,omitempty"`
	// 当前时间戳 (秒)
	Timestamp string `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	// 加密签名
	Signature string `json:"signature,omitempty" xml:"signature,omitempty"`
	// 加密随机字符串
	NonceStr string `json:"nonce_str,omitempty" xml:"nonce_str,omitempty"`
}
