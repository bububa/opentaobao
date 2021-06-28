package film

// TopRefundOrderStatus 
type TopRefundOrderStatus struct {

    // 退款中，其他状态可详见接口文档
    
    Status   string `json:"status,omitempty" xml:"status,omitempty"`
    

    // tbOrderId
    
    TbOrderId   string `json:"tb_order_id,omitempty" xml:"tb_order_id,omitempty"`
    

    // message
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    

}
