package user

// TaobaoKoubeiTribeOpenVerifyCodeApplyResult 结构体
type TaobaoKoubeiTribeOpenVerifyCodeApplyResult struct {
	// traceId
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 错误信息
	Error string `json:"error,omitempty" xml:"error,omitempty"`
	// 用户信息DTO
	Data *UserInfoDto `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
