package tbtrade

// OrderAssembleResponse 结构体
type OrderAssembleResponse struct {
	// 回传结果List
	OrderGroupResponses []OrderGroupResponse `json:"order_group_responses,omitempty" xml:"order_group_responses>order_group_response,omitempty"`
}
