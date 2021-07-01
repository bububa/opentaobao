package alihouse

// AlibabaAlihouseNewhomeLayoutSyncResult 结构体
type AlibabaAlihouseNewhomeLayoutSyncResult struct {
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回户型ID
	Data string `json:"data,omitempty" xml:"data,omitempty"`
}
