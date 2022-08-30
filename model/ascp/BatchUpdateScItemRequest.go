package ascp

// BatchUpdateScItemRequest 结构体
type BatchUpdateScItemRequest struct {
	// 货品列表，最多30条
	ScItems []ScItem `json:"sc_items,omitempty" xml:"sc_items>sc_item,omitempty"`
	// 业务请求ID，用于幂等
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 业务请求时间
	RequestTime int64 `json:"request_time,omitempty" xml:"request_time,omitempty"`
}
