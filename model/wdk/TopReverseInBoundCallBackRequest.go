package wdk

// TopReverseInBoundCallBackRequest 结构体
type TopReverseInBoundCallBackRequest struct {
	// 消息体
	Details []ReverseInBoundDetailCallBackRequest `json:"details,omitempty" xml:"details>reverse_in_bound_detail_call_back_request,omitempty"`
	// 扩展字段
	Extension string `json:"extension,omitempty" xml:"extension,omitempty"`
	// 退仓来源单号/退款单号
	ReverseSourceOrderNo string `json:"reverse_source_order_no,omitempty" xml:"reverse_source_order_no,omitempty"`
	// 仓编码
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
	// 测试标记
	Test bool `json:"test,omitempty" xml:"test,omitempty"`
}
