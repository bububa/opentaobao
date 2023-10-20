package moscm

// AlibabamosorderqueryResultDo 结构体
type AlibabamosorderqueryResultDo struct {
	// 异常信息
	SubMsg string `json:"sub_msg,omitempty" xml:"sub_msg,omitempty"`
	// 等于200代表成功
	SubCode string `json:"sub_code,omitempty" xml:"sub_code,omitempty"`
	// 订单数据
	Data *PagedList `json:"data,omitempty" xml:"data,omitempty"`
	// 调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
