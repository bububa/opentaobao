package logistic

// LogisticsResult 结构体
type LogisticsResult struct {
	// 错误编码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 数据
	Data *Pagination `json:"data,omitempty" xml:"data,omitempty"`
	// 失败消息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
