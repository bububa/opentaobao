package trade

// OrderRefundActionResponse 
/* model for simplify = false
type OrderRefundActionResponse struct {

    // 该订单支持的退款退货操作的集合
    
    SupportRefundActions  struct {
        Supportrefundactions  []Supportrefundactions `json:"supportrefundactions,omitempty"`
    } `json:"support_refund_actions,omitempty"`
    

    // 操作用户ID
    
    OperationUserId   string `json:"operation_user_id,omitempty"`
    

    // 子订单ID
    
    OrderId   string `json:"order_id,omitempty"`
    

}
*/

// OrderRefundActionResponse 
type OrderRefundActionResponse struct {

    // 该订单支持的退款退货操作的集合
    SupportRefundActions   []Supportrefundactions `json:"support_refund_actions,omitempty"`

    // 操作用户ID
    OperationUserId   string `json:"operation_user_id,omitempty"`

    // 子订单ID
    OrderId   string `json:"order_id,omitempty"`

}
