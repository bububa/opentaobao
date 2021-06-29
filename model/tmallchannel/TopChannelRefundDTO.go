package tmallchannel

// TopChannelRefundDTO 
type TopChannelRefundDTO struct {

    // 退款单号
    
    RefundId   int64 `json:"refund_id,omitempty" xml:"refund_id,omitempty"`
    

    // 分销商昵称
    
    DownAccountNick   string `json:"down_account_nick,omitempty" xml:"down_account_nick,omitempty"`
    

    // 服务类型
    
    RefundTypeDesc   string `json:"refund_type_desc,omitempty" xml:"refund_type_desc,omitempty"`
    

    // 单据状态
    
    RefundStatusDesc   string `json:"refund_status_desc,omitempty" xml:"refund_status_desc,omitempty"`
    

    // 采购单号(渠道单)
    
    MainChannelOrderNo   string `json:"main_channel_order_no,omitempty" xml:"main_channel_order_no,omitempty"`
    

    // 交易金额：分
    
    PurchaseOrderPayFee   int64 `json:"purchase_order_pay_fee,omitempty" xml:"purchase_order_pay_fee,omitempty"`
    

    // 退款金额：分
    
    RefundFee   int64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
    

    // 退款原因
    
    RefundReasonDesc   string `json:"refund_reason_desc,omitempty" xml:"refund_reason_desc,omitempty"`
    

    // 退款单创建时间
    
    RefundCreateDate   string `json:"refund_create_date,omitempty" xml:"refund_create_date,omitempty"`
    

    // 铺货状态
    
    PurchaseOrderStatus   string `json:"purchase_order_status,omitempty" xml:"purchase_order_status,omitempty"`
    

    // 销售主订单号（采购单）
    
    MainPurchaseOrderNo   int64 `json:"main_purchase_order_no,omitempty" xml:"main_purchase_order_no,omitempty"`
    

    // 销售子订单号（采购单）
    
    SubPurchaseOrderNo   int64 `json:"sub_purchase_order_no,omitempty" xml:"sub_purchase_order_no,omitempty"`
    

}
