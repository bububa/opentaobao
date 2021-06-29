package uscesl

// TaobaoUsceslBizLightUpResult 
type TaobaoUsceslBizLightUpResult struct {
    // 返回执行码，>=0表示成功
    ReturnCode   int64 `json:"return_code,omitempty" xml:"return_code,omitempty"`
    // 错误编码
    BusinessCode   string `json:"business_code,omitempty" xml:"business_code,omitempty"`
    // 错误描述
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    // 执行结果true或者false
    Target   string `json:"target,omitempty" xml:"target,omitempty"`
    // true或者false
    IsSuccess   string `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
