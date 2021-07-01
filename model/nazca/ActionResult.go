package nazca

// ActionResult 
type ActionResult struct {
    // error
    Error   string `json:"error,omitempty" xml:"error,omitempty"`
    // message
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    // retValue
    RetValue   *AuthApplyDoneCallBackDo `json:"ret_value,omitempty" xml:"ret_value,omitempty"`
    // 错误信息
    SubErrorMessage   string `json:"sub_error_message,omitempty" xml:"sub_error_message,omitempty"`
    // 成功状态
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}
