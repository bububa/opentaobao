package lstvending

// AlibabaLstVendngImageUploadResultDto 结构体
type AlibabaLstVendngImageUploadResultDto struct {
	// 错误信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 图片上传信息
	Module *VendingImageDto `json:"module,omitempty" xml:"module,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 是否处理成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
