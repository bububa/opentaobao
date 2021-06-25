package tvpay

// QueryAuthResultDo 
type QueryAuthResultDo struct {

    // 支付宝账号
    AlipayAccount   string `json:"alipay_account,omitempty"`

    // 授权限额
    LimitAmount   int64 `json:"limit_amount,omitempty"`

    // 授权状态
    Status   string `json:"status,omitempty"`

}
