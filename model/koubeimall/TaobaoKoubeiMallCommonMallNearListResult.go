package koubeimall

// TaobaoKoubeiMallCommonMallNearListResult 结构体
type TaobaoKoubeiMallCommonMallNearListResult struct {
	// 附近商圈列表模型
	MallList []MallDto `json:"mall_list,omitempty" xml:"mall_list>mall_dto,omitempty"`
	// API请求全链路追踪ID
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 错误信息，success=false时，返回相关错误信息
	Error *TribeError `json:"error,omitempty" xml:"error,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
