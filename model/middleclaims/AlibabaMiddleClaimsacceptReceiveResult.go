package middleclaims

// AlibabaMiddleClaimsacceptReceiveResult 
type AlibabaMiddleClaimsacceptReceiveResult struct {
    // 系统调用结果
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 业务结果
    Data   bool `json:"data,omitempty" xml:"data,omitempty"`
    // 是否重复
    Repeated   bool `json:"repeated,omitempty" xml:"repeated,omitempty"`
    // 是否重试
    Retry   bool `json:"retry,omitempty" xml:"retry,omitempty"`
}
