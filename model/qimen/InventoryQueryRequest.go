package qimen

import (
	"sync"
)

// InventoryQueryRequest 结构体
type InventoryQueryRequest struct {
	// 查询准则
	CriteriaList []Criteria `json:"criteriaList,omitempty" xml:"criteriaList>criteria,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 扩展属性
	ExtendProps *TaobaoQimenInventoryQueryMap `json:"extendProps,omitempty" xml:"extendProps,omitempty"`
}

var poolInventoryQueryRequest = sync.Pool{
	New: func() any {
		return new(InventoryQueryRequest)
	},
}

// GetInventoryQueryRequest() 从对象池中获取InventoryQueryRequest
func GetInventoryQueryRequest() *InventoryQueryRequest {
	return poolInventoryQueryRequest.Get().(*InventoryQueryRequest)
}

// ReleaseInventoryQueryRequest 释放InventoryQueryRequest
func ReleaseInventoryQueryRequest(v *InventoryQueryRequest) {
	v.CriteriaList = v.CriteriaList[:0]
	v.Remark = ""
	v.ExtendProps = nil
	poolInventoryQueryRequest.Put(v)
}
