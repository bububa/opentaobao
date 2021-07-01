package moscm

// AlibabaMosOrderRefundListGetResultDto 结构体
type AlibabaMosOrderRefundListGetResultDto struct {
	// 结果集
	Data *PagedList `json:"data,omitempty" xml:"data,omitempty"`
	// 状态码
	SubCode string `json:"sub_code,omitempty" xml:"sub_code,omitempty"`
	// 信息
	SubMsg string `json:"sub_msg,omitempty" xml:"sub_msg,omitempty"`
	// 成功标志
	Success string `json:"success,omitempty" xml:"success,omitempty"`
}
