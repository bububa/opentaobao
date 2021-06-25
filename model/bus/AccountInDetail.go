package bus

// AccountInDetail 
type AccountInDetail struct {

    // 支付宝账号
    AlipayAccount   string `json:"alipay_account,omitempty"`

    // 单位分
    Amount   int64 `json:"amount,omitempty"`

    // 对应该支付宝的支付宝账号ID，注意和支付宝账号保持一致
    AlipayAccountId   string `json:"alipay_account_id,omitempty"`

}
