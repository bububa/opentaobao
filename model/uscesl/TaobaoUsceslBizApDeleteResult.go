package uscesl

// TaobaoUsceslBizApDeleteResult 
type TaobaoUsceslBizApDeleteResult struct {
    // 返回执行对象
    Target   bool `json:"target,omitempty" xml:"target,omitempty"`
    // message
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    // 是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 业务执行返回码
    BusinessCode   string `json:"business_code,omitempty" xml:"business_code,omitempty"`
    // 返回码
    ReturnCode   int64 `json:"return_code,omitempty" xml:"return_code,omitempty"`
}
