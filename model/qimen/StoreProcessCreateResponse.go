package qimen

// StoreProcessCreateResponse 
/* model for simplify = false
type StoreProcessCreateResponse struct {

    // 响应结果:success|failure
    
    Flag   string `json:"flag,omitempty"`
    

    // 响应码
    
    Code   string `json:"code,omitempty"`
    

    // 响应信息
    
    Message   string `json:"message,omitempty"`
    

    // 仓储系统处理单ID
    
    ProcessOrderId   string `json:"processOrderId,omitempty"`
    

}
*/

// StoreProcessCreateResponse 
type StoreProcessCreateResponse struct {

    // 响应结果:success|failure
    Flag   string `json:"flag,omitempty"`

    // 响应码
    Code   string `json:"code,omitempty"`

    // 响应信息
    Message   string `json:"message,omitempty"`

    // 仓储系统处理单ID
    ProcessOrderId   string `json:"processOrderId,omitempty"`

}
