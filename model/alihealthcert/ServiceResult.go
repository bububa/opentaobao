package alihealthcert

// ServiceResult 
type ServiceResult struct {
    // 返回数据对象
    Data   *ReserveStatusResultResponse `json:"data,omitempty" xml:"data,omitempty"`
    // success
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 111
    EagleEyeTraceId   string `json:"eagle_eye_trace_id,omitempty" xml:"eagle_eye_trace_id,omitempty"`
    // 系统错误码
    ErrCode   string `json:"err_code,omitempty" xml:"err_code,omitempty"`
    // 系统错误信息
    ErrMessage   string `json:"err_message,omitempty" xml:"err_message,omitempty"`
}