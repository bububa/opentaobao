package wlb

import (
	"sync"
)

// OutEntityItem 结构体
type OutEntityItem struct {
	// 外部实体类型：&lt;br/&gt;IC_ITEM--ic商品&lt;br/&gt;IC_SKU--ic销售单元
	EntityType string `json:"entity_type,omitempty" xml:"entity_type,omitempty"`
	// entity_type对应的实体类型的id：&lt;br/&gt;当entity_type为IC_ITEM时，entity_id为ic的商品id
	EntityId string `json:"entity_id,omitempty" xml:"entity_id,omitempty"`
}

var poolOutEntityItem = sync.Pool{
	New: func() any {
		return new(OutEntityItem)
	},
}

// GetOutEntityItem() 从对象池中获取OutEntityItem
func GetOutEntityItem() *OutEntityItem {
	return poolOutEntityItem.Get().(*OutEntityItem)
}

// ReleaseOutEntityItem 释放OutEntityItem
func ReleaseOutEntityItem(v *OutEntityItem) {
	v.EntityType = ""
	v.EntityId = ""
	poolOutEntityItem.Put(v)
}
