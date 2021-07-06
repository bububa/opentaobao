package perfect

// BaseResult 结构体
type BaseResult struct {
	// 返回的执行状态吗
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回的数据实体
	Data *PerfectPerformanceItemPublishResp `json:"data,omitempty" xml:"data,omitempty"`
	// 是否执行成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
