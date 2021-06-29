package maitix

// DisEncrypt4CmbParam 
type DisEncrypt4CmbParam struct {

    // 订单金额
    
    Amount   string `json:"amount,omitempty" xml:"amount,omitempty"`
    

    // 订单日期
    
    Date   string `json:"date,omitempty" xml:"date,omitempty"`
    

    // 过期时间跨度
    
    ExpireTimeSpan   string `json:"expire_time_span,omitempty" xml:"expire_time_span,omitempty"`
    

    // 订单号
    
    OrderNo   string `json:"order_no,omitempty" xml:"order_no,omitempty"`
    

    // 支付成功通知的参数,这里传的是大麦的订单号
    
    PayNoticePara   string `json:"pay_notice_para,omitempty" xml:"pay_notice_para,omitempty"`
    

    // 支付成功返回的地址
    
    ReturnUrl   string `json:"return_url,omitempty" xml:"return_url,omitempty"`
    

}
