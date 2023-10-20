package ihome

// AlibabaihomectomcasemainpicupdateApiResult 结构体
type AlibabaihomectomcasemainpicupdateApiResult struct {
	// 具体错误信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// case的url地址
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// true
	Status bool `json:"status,omitempty" xml:"status,omitempty"`
}
