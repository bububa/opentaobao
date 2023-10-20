package alihouse

// AlibabaAlihouseNewhomeProjectUpdateItemInfoResult 结构体
type AlibabaAlihouseNewhomeProjectUpdateItemInfoResult struct {
	// 成功描述
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 成功code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 楼盘表的主键ID
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// 成功返回值
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
