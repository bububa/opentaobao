package tblogistics

// SelectedResourceTopDto 结构体
type SelectedResourceTopDto struct {
	// 下单凭证
	ResourceRequestId string `json:"resource_request_id,omitempty" xml:"resource_request_id,omitempty"`
	// 资源CODE，用于发货
	ResourceCode string `json:"resource_code,omitempty" xml:"resource_code,omitempty"`
	// 实付价格
	ActualFee int64 `json:"actual_fee,omitempty" xml:"actual_fee,omitempty"`
}
