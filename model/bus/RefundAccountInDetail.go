package bus

// RefundAccountInDetail 
type RefundAccountInDetail struct {

    // 收款账户支付宝ID 必传
    AlipayAccount   string `json:"alipay_account,omitempty"`

    // 收款装好支付宝账号，注意和上面的支付宝ID 要对应好
    AlipayAccountId   string `json:"alipay_account_id,omitempty"`

    // 分为单位;退多少钱
    RefundAmount   int64 `json:"refund_amount,omitempty"`

    // 退款批次号须唯一
    RefundBatchNo   string `json:"refund_batch_no,omitempty"`

}
