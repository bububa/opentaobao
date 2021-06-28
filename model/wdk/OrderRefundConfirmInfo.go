package wdk

// OrderRefundConfirmInfo 
/* model for simplify = false
type OrderRefundConfirmInfo struct {

    // 经营店ID
    
    StoreId   string `json:"store_id,omitempty"`
    

    // 渠道店ID
    
    ShopId   string `json:"shop_id,omitempty"`
    

    // 盒马主单号
    
    BizOrderId   string `json:"biz_order_id,omitempty"`
    

    // 同意退款子单
    
    AgreeSubOrders  struct {
        SubRefundConfirm  []SubRefundConfirm `json:"sub_refund_confirm,omitempty"`
    } `json:"agree_sub_orders,omitempty"`
    

    // 外部退款批次Id
    
    OutRefundBatchId   string `json:"out_refund_batch_id,omitempty"`
    

}
*/

// OrderRefundConfirmInfo 
type OrderRefundConfirmInfo struct {

    // 经营店ID
    StoreId   string `json:"store_id,omitempty"`

    // 渠道店ID
    ShopId   string `json:"shop_id,omitempty"`

    // 盒马主单号
    BizOrderId   string `json:"biz_order_id,omitempty"`

    // 同意退款子单
    AgreeSubOrders   []SubRefundConfirm `json:"agree_sub_orders,omitempty"`

    // 外部退款批次Id
    OutRefundBatchId   string `json:"out_refund_batch_id,omitempty"`

}
