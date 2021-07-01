package uscesl

// TaobaoUsceslBizApAddResult 结构体
type TaobaoUsceslBizApAddResult struct {
	// 本次执行返回code
	ReturnCode int64 `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// 业务错误code
	BusinessCode string `json:"business_code,omitempty" xml:"business_code,omitempty"`
	// 执行成功或失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 业务错误文案
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回本次执行成功或失败
	Target bool `json:"target,omitempty" xml:"target,omitempty"`
}
