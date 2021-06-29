package koubeimall

// TribeError 
type TribeError struct {
    // 错误信息可读性描述
    View   string `json:"view,omitempty" xml:"view,omitempty"`
    // 错误信息描述
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    // 错误码
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
}
