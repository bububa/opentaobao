package normalvisa

// TaobaoAlitripTravelNormalvisaUploadfileResultSet 结构体
type TaobaoAlitripTravelNormalvisaUploadfileResultSet struct {
	// 结果列表
	Results []bool `json:"results,omitempty" xml:"results>bool,omitempty"`
	// 异常
	Exception string `json:"exception,omitempty" xml:"exception,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 数量
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
	// 是否包含下一个
	HasNext bool `json:"has_next,omitempty" xml:"has_next,omitempty"`
}
