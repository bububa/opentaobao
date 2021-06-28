package alsc

// RefundInfo 
type RefundInfo struct {

    // 退款金额
    
    RefundAmount   int64 `json:"refund_amount,omitempty" xml:"refund_amount,omitempty"`
    

    // 退款商品
    
    RefundItemList   []Item `json:"refund_item_list,omitempty" xml:"refund_item_list,omitempty"`
    

    // 退款流水号（支付宝单号、微信单号、三方单号）
    
    RefundPayNo   string `json:"refund_pay_no,omitempty" xml:"refund_pay_no,omitempty"`
    

    // 退款状态：  INITIAL("INITIAL", "退款初始状态"),  PROCESSING("PROCESSING", "处理中状态"),  SUCCESS("SUCCESS", "退款成功状态"),  FAIL("FAIL", "退款失败状态");
    
    RefundStatus   string `json:"refund_status,omitempty" xml:"refund_status,omitempty"`
    

    // 退款时间
    
    RefundTime   string `json:"refund_time,omitempty" xml:"refund_time,omitempty"`
    

    // 退款类型 ：  整单-ALL  部分-PART
    
    RefundType   string `json:"refund_type,omitempty" xml:"refund_type,omitempty"`
    

    // 退款单号，必填
    
    OutRefundNo   string `json:"out_refund_no,omitempty" xml:"out_refund_no,omitempty"`
    

}
