package baoxian

// UploadResult 结构体
type UploadResult struct {
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// errorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// model
	Model *InsAttachmentUploadVo `json:"model,omitempty" xml:"model,omitempty"`
	// isSuccess
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
