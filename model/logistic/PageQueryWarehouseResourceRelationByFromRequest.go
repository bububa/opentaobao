package logistic

import (
	"sync"
)

// PageQueryWarehouseResourceRelationByFromRequest 结构体
type PageQueryWarehouseResourceRelationByFromRequest struct {
	// from资源外部编码,与from_resource_code二选一
	FromOrgResourceCode string `json:"from_org_resource_code,omitempty" xml:"from_org_resource_code,omitempty"`
	// from资源编码，与from_org_resource_code二选一
	FromResourceCode string `json:"from_resource_code,omitempty" xml:"from_resource_code,omitempty"`
	// from资源类型
	FromResourceType string `json:"from_resource_type,omitempty" xml:"from_resource_type,omitempty"`
	// 网络编码
	NetworkCode string `json:"network_code,omitempty" xml:"network_code,omitempty"`
	// 分页，1开始
	PageIndex int64 `json:"page_index,omitempty" xml:"page_index,omitempty"`
	// 分页，上限50
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}

var poolPageQueryWarehouseResourceRelationByFromRequest = sync.Pool{
	New: func() any {
		return new(PageQueryWarehouseResourceRelationByFromRequest)
	},
}

// GetPageQueryWarehouseResourceRelationByFromRequest() 从对象池中获取PageQueryWarehouseResourceRelationByFromRequest
func GetPageQueryWarehouseResourceRelationByFromRequest() *PageQueryWarehouseResourceRelationByFromRequest {
	return poolPageQueryWarehouseResourceRelationByFromRequest.Get().(*PageQueryWarehouseResourceRelationByFromRequest)
}

// ReleasePageQueryWarehouseResourceRelationByFromRequest 释放PageQueryWarehouseResourceRelationByFromRequest
func ReleasePageQueryWarehouseResourceRelationByFromRequest(v *PageQueryWarehouseResourceRelationByFromRequest) {
	v.FromOrgResourceCode = ""
	v.FromResourceCode = ""
	v.FromResourceType = ""
	v.NetworkCode = ""
	v.PageIndex = 0
	v.PageSize = 0
	poolPageQueryWarehouseResourceRelationByFromRequest.Put(v)
}
