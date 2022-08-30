package ascp

// ItemMappingUpdateAsyncRequest 结构体
type ItemMappingUpdateAsyncRequest struct {
	// 商货品关联列表
	ItemMappings []ItemMapping `json:"item_mappings,omitempty" xml:"item_mappings>item_mapping,omitempty"`
	// 业务请求ID，用于做幂等
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// ERP调翱象接口创建货品的时间戳
	RequestTime int64 `json:"request_time,omitempty" xml:"request_time,omitempty"`
}
