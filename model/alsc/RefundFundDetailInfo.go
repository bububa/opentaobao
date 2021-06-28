package alsc

// RefundFundDetailInfo 
type RefundFundDetailInfo struct {

    // 外部订单号
    
    OutOrderNo   string `json:"out_order_no,omitempty" xml:"out_order_no,omitempty"`
    

    // 支付方式id
    
    OutPayChannelId   string `json:"out_pay_channel_id,omitempty" xml:"out_pay_channel_id,omitempty"`
    

    // 支付方式名称
    
    OutPayChannelName   string `json:"out_pay_channel_name,omitempty" xml:"out_pay_channel_name,omitempty"`
    

    // 支付方式code
    
    OutPayCode   string `json:"out_pay_code,omitempty" xml:"out_pay_code,omitempty"`
    

    // 外部退款单资金明细id
    
    OutRefundFundDetailNo   string `json:"out_refund_fund_detail_no,omitempty" xml:"out_refund_fund_detail_no,omitempty"`
    

    // 外部退款单号，支付宝，微信等退款单号
    
    OutRefundNo   string `json:"out_refund_no,omitempty" xml:"out_refund_no,omitempty"`
    

    // 外部退款单id
    
    OutRefundOrderNo   string `json:"out_refund_order_no,omitempty" xml:"out_refund_order_no,omitempty"`
    

    // 退款金额
    
    RefundFee   int64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
    

    // SUCCESS成功  FAIL失败
    
    Status   string `json:"status,omitempty" xml:"status,omitempty"`
    

    // 退款成功时间
    
    SuccessTime   int64 `json:"success_time,omitempty" xml:"success_time,omitempty"`
    

    // 扩展信息
    
    ExtInfo   string `json:"ext_info,omitempty" xml:"ext_info,omitempty"`
    

}
