package wdk

// TopApiResult 结构体
type TopApiResult struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误详情
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 数据
	Data *WarehouseNetworkResultDto `json:"data,omitempty" xml:"data,omitempty"`
	// 请求成功或失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
