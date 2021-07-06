package tmallcar

// TmallAliautoWisdomdataModelReciveResult 结构体
type TmallAliautoWisdomdataModelReciveResult struct {
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回的执行状态吗
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
