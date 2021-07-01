package moscm

// AlibabaMosDeliverySendResultDo 
type AlibabaMosDeliverySendResultDo struct {
    // 异常信息
    SubMsg   string `json:"sub_msg,omitempty" xml:"sub_msg,omitempty"`
    // 编码
    SubCode   string `json:"sub_code,omitempty" xml:"sub_code,omitempty"`
    // 调用是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}
