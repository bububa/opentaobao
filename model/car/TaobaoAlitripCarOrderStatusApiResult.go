package car

// TaobaoalitripcarorderstatusApiResult 结构体
type TaobaoalitripcarorderstatusApiResult struct {
	// 其它数据
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 状态码
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
}
