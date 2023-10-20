package store

import (
	"sync"
)

// StoreRelationSimpleDo 结构体
type StoreRelationSimpleDo struct {
	// 门店名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 业务id
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 业务类型
	BizType int64 `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
	// 关系id
	RelationId int64 `json:"relation_id,omitempty" xml:"relation_id,omitempty"`
	// 业务关系状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 门店id
	StoreId int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
}

var poolStoreRelationSimpleDo = sync.Pool{
	New: func() any {
		return new(StoreRelationSimpleDo)
	},
}

// GetStoreRelationSimpleDo() 从对象池中获取StoreRelationSimpleDo
func GetStoreRelationSimpleDo() *StoreRelationSimpleDo {
	return poolStoreRelationSimpleDo.Get().(*StoreRelationSimpleDo)
}

// ReleaseStoreRelationSimpleDo 释放StoreRelationSimpleDo
func ReleaseStoreRelationSimpleDo(v *StoreRelationSimpleDo) {
	v.Name = ""
	v.OuterId = ""
	v.BizType = 0
	v.RelationId = 0
	v.Status = 0
	v.StoreId = 0
	poolStoreRelationSimpleDo.Put(v)
}
