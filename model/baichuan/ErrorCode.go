package baichuan

// ErrorCode 
type ErrorCode struct {

    // 详细错误信息
    Message   string `json:"message,omitempty"`

    // 错误码
    Code   string `json:"code,omitempty"`

}
