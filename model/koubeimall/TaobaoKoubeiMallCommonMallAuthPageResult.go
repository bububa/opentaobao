package koubeimall

// TaobaoKoubeiMallCommonMallAuthPageResult 结构体
type TaobaoKoubeiMallCommonMallAuthPageResult struct {
	// API请求全链路追踪ID
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// API返回的data模型
	Data *PageResult `json:"data,omitempty" xml:"data,omitempty"`
	// 请求是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 错误信息，success=false时，返回相关错误信息
	Error *TribeError `json:"error,omitempty" xml:"error,omitempty"`
}
