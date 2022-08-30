package ascp

// BatchUpdateCombineScItemRequest 结构体
type BatchUpdateCombineScItemRequest struct {
	// 组合货品列表，批量数量不可超过30
	CombineScItems []CombineScItem `json:"combine_sc_items,omitempty" xml:"combine_sc_items>combine_sc_item,omitempty"`
	// 业务请求ID，用于幂等
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 业务请求时间
	RequestTime int64 `json:"request_time,omitempty" xml:"request_time,omitempty"`
}
