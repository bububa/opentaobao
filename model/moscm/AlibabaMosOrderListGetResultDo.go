package moscm

// AlibabaMosOrderListGetResultDo 结构体
type AlibabaMosOrderListGetResultDo struct {
	// 消息
	SubMsg string `json:"sub_msg,omitempty" xml:"sub_msg,omitempty"`
	// 状态码
	SubCode string `json:"sub_code,omitempty" xml:"sub_code,omitempty"`
	// 结果集
	Data *PagedList `json:"data,omitempty" xml:"data,omitempty"`
	// 成功标志
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
