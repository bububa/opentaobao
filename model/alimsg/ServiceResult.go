package alimsg

// ServiceResult 
type ServiceResult struct {
    // 内容
    Content   *SendTemplateMsgResponse `json:"content,omitempty" xml:"content,omitempty"`
    // 错误码
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    // 错误码描述
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
    // success
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}
