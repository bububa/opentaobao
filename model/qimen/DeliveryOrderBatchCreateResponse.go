package qimen

// DeliveryOrderBatchCreateResponse 
/* model for simplify = false
type DeliveryOrderBatchCreateResponse struct {

    // 响应结果:success|failure
    
    Flag   string `json:"flag,omitempty"`
    

    // 响应码
    
    Code   string `json:"code,omitempty"`
    

    // 响应信息
    
    Message   string `json:"message,omitempty"`
    

    // 订单详情
    
    Orders  struct {
        Order  []Order `json:"order,omitempty"`
    } `json:"orders,omitempty"`
    

}
*/

// DeliveryOrderBatchCreateResponse 
type DeliveryOrderBatchCreateResponse struct {

    // 响应结果:success|failure
    Flag   string `json:"flag,omitempty"`

    // 响应码
    Code   string `json:"code,omitempty"`

    // 响应信息
    Message   string `json:"message,omitempty"`

    // 订单详情
    Orders   []Order `json:"orders,omitempty"`

}
