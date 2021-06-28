package alicom

// CommonResult 
/* model for simplify = false
type CommonResult struct {

    // 接口返回成功
    
    Success   bool `json:"success,omitempty"`
    

    // code
    
    Code   string `json:"code,omitempty"`
    

    // desc
    
    Desc   string `json:"desc,omitempty"`
    

    // 阿里订单号
    
    Model   string `json:"model,omitempty"`
    

}
*/

// CommonResult 
type CommonResult struct {

    // 接口返回成功
    Success   bool `json:"success,omitempty"`

    // code
    Code   string `json:"code,omitempty"`

    // desc
    Desc   string `json:"desc,omitempty"`

    // 阿里订单号
    Model   string `json:"model,omitempty"`

}
