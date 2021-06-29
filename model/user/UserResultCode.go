package user

// UserResultCode 
type UserResultCode struct {
    // 结果code
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`
    // 结果类型
    ResultType   string `json:"result_type,omitempty" xml:"result_type,omitempty"`
    // 结果描述
    ResultMsg   string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
}
