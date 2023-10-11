package wdk

// OrderDesensitizePhoneResult 结构体
type OrderDesensitizePhoneResult struct {
	// 商家编码
	MerchantCode string `json:"merchant_code,omitempty" xml:"merchant_code,omitempty"`
	// 业务单号
	BizOrderId string `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// 经营店编码
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
	// 脱敏的手机号
	ReceiverPrivacyPhone string `json:"receiver_privacy_phone,omitempty" xml:"receiver_privacy_phone,omitempty"`
	// 虚拟号
	VirtualNumber string `json:"virtual_number,omitempty" xml:"virtual_number,omitempty"`
	// 买家虚拟号（仅鲜花行业订单有值）
	BuyerVirtualNumber string `json:"buyer_virtual_number,omitempty" xml:"buyer_virtual_number,omitempty"`
	// 预计过期时间
	ExpirationTime int64 `json:"expiration_time,omitempty" xml:"expiration_time,omitempty"`
}
