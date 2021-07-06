package tmallgenie

// AlibabaAilabsIotDeviceListGetResult 结构体
type AlibabaAilabsIotDeviceListGetResult struct {
	// 返回值list
	RetValues []RetValue `json:"ret_values,omitempty" xml:"ret_values>ret_value,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误码
	RetCode int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
