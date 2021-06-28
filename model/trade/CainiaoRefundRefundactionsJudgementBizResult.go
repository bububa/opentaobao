package trade

// CainiaoRefundRefundactionsJudgementBizResult 
/* model for simplify = false
type CainiaoRefundRefundactionsJudgementBizResult struct {

    // 调用时错误码
    
    StatusCode   string `json:"status_code,omitempty"`
    

    // 是否能对某一个订单进行退货退款操作的返回值
    
    Data  *struct {
        OrderRefundOperationResponse  *OrderRefundOperationResponse `json:"order_refund_operation_response,omitempty"`
    } `json:"data,omitempty"`
    

    // true表示成功，false表示失败
    
    Success   bool `json:"success,omitempty"`
    

    // 调用时错误描述
    
    StatusMessage   string `json:"status_message,omitempty"`
    

}
*/

// CainiaoRefundRefundactionsJudgementBizResult 
type CainiaoRefundRefundactionsJudgementBizResult struct {

    // 调用时错误码
    StatusCode   string `json:"status_code,omitempty"`

    // 是否能对某一个订单进行退货退款操作的返回值
    Data   *OrderRefundOperationResponse `json:"data,omitempty"`

    // true表示成功，false表示失败
    Success   bool `json:"success,omitempty"`

    // 调用时错误描述
    StatusMessage   string `json:"status_message,omitempty"`

}
