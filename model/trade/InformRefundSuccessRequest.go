package trade

// InformRefundSuccessRequest 
/* model for simplify = false
type InformRefundSuccessRequest struct {

    // 门店编码
    
    ShopId   string `json:"shop_id,omitempty"`
    

    // 退货成功履约子单
    
    RefundFulfillSubOrders  struct {
        RefundSuccessSubOrderRequest  []RefundSuccessSubOrderRequest `json:"refund_success_sub_order_request,omitempty"`
    } `json:"refund_fulfill_sub_orders,omitempty"`
    

}
*/

// InformRefundSuccessRequest 
type InformRefundSuccessRequest struct {

    // 门店编码
    ShopId   string `json:"shop_id,omitempty"`

    // 退货成功履约子单
    RefundFulfillSubOrders   []RefundSuccessSubOrderRequest `json:"refund_fulfill_sub_orders,omitempty"`

}
