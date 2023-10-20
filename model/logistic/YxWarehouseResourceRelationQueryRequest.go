package logistic

import (
	"sync"
)

// YxWarehouseResourceRelationQueryRequest 结构体
type YxWarehouseResourceRelationQueryRequest struct {
	// 网格仓外部编码
	ToOrgResourceCodes []string `json:"to_org_resource_codes,omitempty" xml:"to_org_resource_codes>string,omitempty"`
}

var poolYxWarehouseResourceRelationQueryRequest = sync.Pool{
	New: func() any {
		return new(YxWarehouseResourceRelationQueryRequest)
	},
}

// GetYxWarehouseResourceRelationQueryRequest() 从对象池中获取YxWarehouseResourceRelationQueryRequest
func GetYxWarehouseResourceRelationQueryRequest() *YxWarehouseResourceRelationQueryRequest {
	return poolYxWarehouseResourceRelationQueryRequest.Get().(*YxWarehouseResourceRelationQueryRequest)
}

// ReleaseYxWarehouseResourceRelationQueryRequest 释放YxWarehouseResourceRelationQueryRequest
func ReleaseYxWarehouseResourceRelationQueryRequest(v *YxWarehouseResourceRelationQueryRequest) {
	v.ToOrgResourceCodes = v.ToOrgResourceCodes[:0]
	poolYxWarehouseResourceRelationQueryRequest.Put(v)
}
