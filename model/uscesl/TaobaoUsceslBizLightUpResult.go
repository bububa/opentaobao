package uscesl

// TaobaoUsceslBizLightUpResult 结构体
type TaobaoUsceslBizLightUpResult struct {
	// 执行结果true或者false
	Target string `json:"target,omitempty" xml:"target,omitempty"`
	// 错误编码
	BusinessCode string `json:"business_code,omitempty" xml:"business_code,omitempty"`
	// 错误描述
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// true或者false
	IsSuccess string `json:"is_success,omitempty" xml:"is_success,omitempty"`
	// 返回执行码，&gt;=0表示成功
	ReturnCode int64 `json:"return_code,omitempty" xml:"return_code,omitempty"`
}
