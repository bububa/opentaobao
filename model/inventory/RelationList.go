package inventory

import (
	"sync"
)

// RelationList 结构体
type RelationList struct {
	// 生效的前端宝贝id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 生效的前端宝贝的skuid
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
}

var poolRelationList = sync.Pool{
	New: func() any {
		return new(RelationList)
	},
}

// GetRelationList() 从对象池中获取RelationList
func GetRelationList() *RelationList {
	return poolRelationList.Get().(*RelationList)
}

// ReleaseRelationList 释放RelationList
func ReleaseRelationList(v *RelationList) {
	v.ItemId = 0
	v.SkuId = 0
	poolRelationList.Put(v)
}
