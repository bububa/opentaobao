package wdk

// OrderRefundConfirmInfo 
type OrderRefundConfirmInfo struct {

    // 经营店ID
    
    StoreId   string `json:"store_id,omitempty" xml:"store_id,omitempty"`
    

    // 渠道店ID
    
    ShopId   string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
    

    // 盒马主单号
    
    BizOrderId   string `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
    

    // 同意退款子单
    
    AgreeSubOrders   []SubRefundConfirm `json:"agree_sub_orders,omitempty" xml:"agree_sub_orders,omitempty"`
    

    // 外部退款批次Id
    
    OutRefundBatchId   string `json:"out_refund_batch_id,omitempty" xml:"out_refund_batch_id,omitempty"`
    

}
