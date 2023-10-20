package ascp

import (
	"sync"
)

// ItemMappingDeleteRequest 结构体
type ItemMappingDeleteRequest struct {
	// 商货品关联关系
	ItemMappings []ItemMapping `json:"item_mappings,omitempty" xml:"item_mappings>item_mapping,omitempty"`
	// 业务请求ID，用于做幂等
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// ERP调翱象接口创建货品的时间戳
	RequestTime int64 `json:"request_time,omitempty" xml:"request_time,omitempty"`
}

var poolItemMappingDeleteRequest = sync.Pool{
	New: func() any {
		return new(ItemMappingDeleteRequest)
	},
}

// GetItemMappingDeleteRequest() 从对象池中获取ItemMappingDeleteRequest
func GetItemMappingDeleteRequest() *ItemMappingDeleteRequest {
	return poolItemMappingDeleteRequest.Get().(*ItemMappingDeleteRequest)
}

// ReleaseItemMappingDeleteRequest 释放ItemMappingDeleteRequest
func ReleaseItemMappingDeleteRequest(v *ItemMappingDeleteRequest) {
	v.ItemMappings = v.ItemMappings[:0]
	v.RequestId = ""
	v.RequestTime = 0
	poolItemMappingDeleteRequest.Put(v)
}
