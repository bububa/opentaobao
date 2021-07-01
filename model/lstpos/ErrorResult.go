package lstpos

// ErrorResult 结构体
type ErrorResult struct {
	// 单个订单错误消息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 单个订单错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 单个订单isv业务处理关键字值
	Key string `json:"key,omitempty" xml:"key,omitempty"`
	// 单个订单处理结果标示  true：成功 false：失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 业务处理数据
	Data string `json:"data,omitempty" xml:"data,omitempty"`
}
