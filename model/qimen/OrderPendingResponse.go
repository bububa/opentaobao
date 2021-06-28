package qimen

// OrderPendingResponse 
/* model for simplify = false
type OrderPendingResponse struct {

    // 响应结果:success|failure
    
    Flag   string `json:"flag,omitempty"`
    

    // 响应码
    
    Code   string `json:"code,omitempty"`
    

    // 响应信息
    
    Message   string `json:"message,omitempty"`
    

}
*/

// OrderPendingResponse 
type OrderPendingResponse struct {

    // 响应结果:success|failure
    Flag   string `json:"flag,omitempty"`

    // 响应码
    Code   string `json:"code,omitempty"`

    // 响应信息
    Message   string `json:"message,omitempty"`

}
