package alicom

// CommonResult 
type CommonResult struct {
    // 接口返回成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // code
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    // desc
    Desc   string `json:"desc,omitempty" xml:"desc,omitempty"`
    // 阿里订单号
    Model   string `json:"model,omitempty" xml:"model,omitempty"`
}
