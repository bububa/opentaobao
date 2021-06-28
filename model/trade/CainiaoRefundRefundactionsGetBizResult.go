package trade

// CainiaoRefundRefundactionsGetBizResult 
/* model for simplify = false
type CainiaoRefundRefundactionsGetBizResult struct {

    // 调用错误时，错误code
    
    StatusCode   string `json:"status_code,omitempty"`
    

    // 该订单支持的退款退货操作的集合
    
    Data  *struct {
        OrderRefundActionResponse  *OrderRefundActionResponse `json:"order_refund_action_response,omitempty"`
    } `json:"data,omitempty"`
    

    // true表示成功，false表示失败
    
    Success   bool `json:"success,omitempty"`
    

    // 调用错误时，错误描述
    
    StatusMessage   string `json:"status_message,omitempty"`
    

}
*/

// CainiaoRefundRefundactionsGetBizResult 
type CainiaoRefundRefundactionsGetBizResult struct {

    // 调用错误时，错误code
    StatusCode   string `json:"status_code,omitempty"`

    // 该订单支持的退款退货操作的集合
    Data   *OrderRefundActionResponse `json:"data,omitempty"`

    // true表示成功，false表示失败
    Success   bool `json:"success,omitempty"`

    // 调用错误时，错误描述
    StatusMessage   string `json:"status_message,omitempty"`

}
