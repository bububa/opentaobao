package wdk

// RefundCancelInfo 
/* model for simplify = false
type RefundCancelInfo struct {

    // 外部逆向单ID
    
    OutRefundId   string `json:"out_refund_id,omitempty"`
    

    // 外部主单号
    
    OutOrderId   string `json:"out_order_id,omitempty"`
    

    // 外部渠道店ID
    
    OutShopId   string `json:"out_shop_id,omitempty"`
    

    // 订单来源
    
    OrderFrom   int64 `json:"order_from,omitempty"`
    

}
*/

// RefundCancelInfo 
type RefundCancelInfo struct {

    // 外部逆向单ID
    OutRefundId   string `json:"out_refund_id,omitempty"`

    // 外部主单号
    OutOrderId   string `json:"out_order_id,omitempty"`

    // 外部渠道店ID
    OutShopId   string `json:"out_shop_id,omitempty"`

    // 订单来源
    OrderFrom   int64 `json:"order_from,omitempty"`

}
