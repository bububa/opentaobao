package ascp

// BatchCreateItemMappingRequest 结构体
type BatchCreateItemMappingRequest struct {
	// 商货品关联列表，批量数量不可超过30
	ItemMappings []ItemMapping `json:"item_mappings,omitempty" xml:"item_mappings>item_mapping,omitempty"`
	// 业务请求ID，用于幂等
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 业务请求时间
	RequestTime int64 `json:"request_time,omitempty" xml:"request_time,omitempty"`
}
