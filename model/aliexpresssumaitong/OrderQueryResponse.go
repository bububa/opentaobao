package aliexpresssumaitong

// OrderQueryResponse 
type OrderQueryResponse struct {

    // 错误提示信息
    
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
    

    // 订单信息
    
    Orders   []Order `json:"orders,omitempty" xml:"orders,omitempty"`
    

    // 接口调用结果
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

}
