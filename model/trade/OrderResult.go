package trade

// OrderResult 
/* model for simplify = false
type OrderResult struct {

    // 错误编码
    
    ErrorCode   string `json:"error_code,omitempty"`
    

    // 错误消息
    
    ErrorMsg   string `json:"error_msg,omitempty"`
    

    // 是否取消成功
    
    Success   bool `json:"success,omitempty"`
    

    // 取消的订单
    
    Result  *struct {
        OrderObject  *OrderObject `json:"order_object,omitempty"`
    } `json:"result,omitempty"`
    

    // 创建的订单
    
    Trade  *struct {
        TradeOrder  *TradeOrder `json:"trade_order,omitempty"`
    } `json:"trade,omitempty"`
    

}
*/

// OrderResult 
type OrderResult struct {

    // 错误编码
    ErrorCode   string `json:"error_code,omitempty"`

    // 错误消息
    ErrorMsg   string `json:"error_msg,omitempty"`

    // 是否取消成功
    Success   bool `json:"success,omitempty"`

    // 取消的订单
    Result   *OrderObject `json:"result,omitempty"`

    // 创建的订单
    Trade   *TradeOrder `json:"trade,omitempty"`

}
