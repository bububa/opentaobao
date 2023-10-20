package fpm

// AlibabafpmfileuploadResultModel 结构体
type AlibabafpmfileuploadResultModel struct {
	// 错误编码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错信信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// returnValue
	ReturnValue *FileUploadReponseDto `json:"return_value,omitempty" xml:"return_value,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
