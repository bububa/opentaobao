package bus

// AccountInDetail 结构体
type AccountInDetail struct {
	// 支付宝账号
	AlipayAccount string `json:"alipay_account,omitempty" xml:"alipay_account,omitempty"`
	// 对应该支付宝的支付宝账号ID，注意和支付宝账号保持一致
	AlipayAccountId string `json:"alipay_account_id,omitempty" xml:"alipay_account_id,omitempty"`
	// 单位分
	Amount int64 `json:"amount,omitempty" xml:"amount,omitempty"`
}
