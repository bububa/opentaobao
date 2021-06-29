package lsttrade

// TopLstSupplierOrderRefundQuery 
type TopLstSupplierOrderRefundQuery struct {

    // 1:售中退款，2:售后退款；0:所有退款单
    
    RefundType   int64 `json:"refund_type,omitempty" xml:"refund_type,omitempty"`
    

    // 退款申请时间（起始）
    
    ApplyEndTime   string `json:"apply_end_time,omitempty" xml:"apply_end_time,omitempty"`
    

    // 主订单Id
    
    OrderId   int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
    

    // 查询页码
    
    CurrentPageNum   int64 `json:"current_page_num,omitempty" xml:"current_page_num,omitempty"`
    

    // 页大小
    
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    

    // 退款状态，等待卖家同意("waitselleragree"), 退款成功("refundsuccess"), 退款关闭("refundclose"), 待买家修改("waitbuyermodify"), 等待买家退货("waitbuyersend"), 等待卖家确认收货("waitsellerreceive");
    
    RefundStatus   string `json:"refund_status,omitempty" xml:"refund_status,omitempty"`
    

    // 买家userid
    
    BuyerUserId   int64 `json:"buyer_user_id,omitempty" xml:"buyer_user_id,omitempty"`
    

    // 退款单id
    
    RefundId   string `json:"refund_id,omitempty" xml:"refund_id,omitempty"`
    

    // 退款申请时间(截止)
    
    ApplyStartTime   string `json:"apply_start_time,omitempty" xml:"apply_start_time,omitempty"`
    

}
