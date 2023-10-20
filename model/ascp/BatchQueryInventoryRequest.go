package ascp

import (
	"sync"
)

// BatchQueryInventoryRequest 结构体
type BatchQueryInventoryRequest struct {
	// 货品集合
	ScItemList []ScItem `json:"sc_item_list,omitempty" xml:"sc_item_list>sc_item,omitempty"`
	// 业务请求ID
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 货主编码；优仓分配，长度不会超过32位，不含特殊字符
	OwnerCode string `json:"owner_code,omitempty" xml:"owner_code,omitempty"`
	// 业务请求时间(毫秒数)
	RequestTime int64 `json:"request_time,omitempty" xml:"request_time,omitempty"`
}

var poolBatchQueryInventoryRequest = sync.Pool{
	New: func() any {
		return new(BatchQueryInventoryRequest)
	},
}

// GetBatchQueryInventoryRequest() 从对象池中获取BatchQueryInventoryRequest
func GetBatchQueryInventoryRequest() *BatchQueryInventoryRequest {
	return poolBatchQueryInventoryRequest.Get().(*BatchQueryInventoryRequest)
}

// ReleaseBatchQueryInventoryRequest 释放BatchQueryInventoryRequest
func ReleaseBatchQueryInventoryRequest(v *BatchQueryInventoryRequest) {
	v.ScItemList = v.ScItemList[:0]
	v.RequestId = ""
	v.OwnerCode = ""
	v.RequestTime = 0
	poolBatchQueryInventoryRequest.Put(v)
}
