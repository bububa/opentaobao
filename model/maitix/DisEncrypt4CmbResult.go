package maitix

// DisEncrypt4CmbResult 
type DisEncrypt4CmbResult struct {
    // 订单金额
    Amount   string `json:"amount,omitempty" xml:"amount,omitempty"`
    // 分行号
    BranchNo   string `json:"branch_no,omitempty" xml:"branch_no,omitempty"`
    // 订单日期
    Date   string `json:"date,omitempty" xml:"date,omitempty"`
    // 请求日间
    DateTime   string `json:"date_time,omitempty" xml:"date_time,omitempty"`
    // 过期时间跨度
    ExpireTimeSpan   string `json:"expire_time_span,omitempty" xml:"expire_time_span,omitempty"`
    // 商户号
    MerchantNo   string `json:"merchant_no,omitempty" xml:"merchant_no,omitempty"`
    // 订单号
    OrderNo   string `json:"order_no,omitempty" xml:"order_no,omitempty"`
    // 支付成功通知的参数,这里传的是大麦的订单号
    PayNoticePara   string `json:"pay_notice_para,omitempty" xml:"pay_notice_para,omitempty"`
    // 支付成功通知的地址
    PayNoticeUrl   string `json:"pay_notice_url,omitempty" xml:"pay_notice_url,omitempty"`
    // 支付成功返回的地址
    ReturnUrl   string `json:"return_url,omitempty" xml:"return_url,omitempty"`
    // 加密结果
    Sign   string `json:"sign,omitempty" xml:"sign,omitempty"`
}
