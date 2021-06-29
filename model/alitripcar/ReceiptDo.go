package alitripcar

// ReceiptDo 
type ReceiptDo struct {

    // 发票金额，单位分
    
    ReceiptAmount   int64 `json:"receipt_amount,omitempty" xml:"receipt_amount,omitempty"`
    

    // 业务类型编号（1：?国内机票，2：国际机票，3：火车票，4：汽车票，5：酒店，6：门票度假，7：打车，8：用车，9：套餐，10：欧铁，11：辅营，12：辅营保险）
    
    BizType   int64 `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
    

    // 错误码
    
    FailCode   string `json:"fail_code,omitempty" xml:"fail_code,omitempty"`
    

    // 失败原因.status为0时必填
    
    FailReason   string `json:"fail_reason,omitempty" xml:"fail_reason,omitempty"`
    

    // 发票链接.status为1时需要填
    
    ReceiptUrl   string `json:"receipt_url,omitempty" xml:"receipt_url,omitempty"`
    

    // 开票时间
    
    ReceiptDateTime   string `json:"receipt_date_time,omitempty" xml:"receipt_date_time,omitempty"`
    

    // 发票流水号。成功时必填
    
    ReceiptNumber   string `json:"receipt_number,omitempty" xml:"receipt_number,omitempty"`
    

    // 发票状态1成功0失败-1取消订单
    
    ReceiptStatus   int64 `json:"receipt_status,omitempty" xml:"receipt_status,omitempty"`
    

    // 供应商编号
    
    ProviderId   string `json:"provider_id,omitempty" xml:"provider_id,omitempty"`
    

    // 飞猪订单号
    
    OrderId   int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
    

    // 供应商订单号
    
    OutOrderId   string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
    

}
