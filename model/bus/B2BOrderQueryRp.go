package bus

// B2BOrderQueryRp 
/* model for simplify = false
type B2BOrderQueryRp struct {

    // 订单对象
    
    B2bBusOrderInfo  *struct {
        B2BBusOrderInfo  *B2BBusOrderInfo `json:"b2b_bus_order_info,omitempty"`
    } `json:"b2b_bus_order_info,omitempty"`
    

    // 错误代码
    
    ErrorCode   string `json:"error_code,omitempty"`
    

    // 错误描述
    
    ErrorMsg   string `json:"error_msg,omitempty"`
    

    // 是否查询成功
    
    Success   bool `json:"success,omitempty"`
    

}
*/

// B2BOrderQueryRp 
type B2BOrderQueryRp struct {

    // 订单对象
    B2bBusOrderInfo   *B2BBusOrderInfo `json:"b2b_bus_order_info,omitempty"`

    // 错误代码
    ErrorCode   string `json:"error_code,omitempty"`

    // 错误描述
    ErrorMsg   string `json:"error_msg,omitempty"`

    // 是否查询成功
    Success   bool `json:"success,omitempty"`

}
