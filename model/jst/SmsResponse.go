package jst

// SmsResponse 
type SmsResponse struct {
    // 系统异常
    ResponseCode   string `json:"response_code,omitempty" xml:"response_code,omitempty"`
    // 成功
    ResponseSuccess   bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
    // 请求id
    RequestId   string `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回值
    Module   *SmsStateModel `json:"module,omitempty" xml:"module,omitempty"`
    // 成功
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    // 请求成功
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    // 请求成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}
