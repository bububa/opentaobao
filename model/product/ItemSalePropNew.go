package product

import (
	"sync"
)

// ItemSalePropNew 结构体
type ItemSalePropNew struct {
	// 属性状态集合
	SalePropValueStatusList []SalePropValueStatus `json:"sale_prop_value_status_list,omitempty" xml:"sale_prop_value_status_list>sale_prop_value_status,omitempty"`
	// 属性文本
	PropertyValue string `json:"property_value,omitempty" xml:"property_value,omitempty"`
	// 属性ID
	PropertyId int64 `json:"property_id,omitempty" xml:"property_id,omitempty"`
}

var poolItemSalePropNew = sync.Pool{
	New: func() any {
		return new(ItemSalePropNew)
	},
}

// GetItemSalePropNew() 从对象池中获取ItemSalePropNew
func GetItemSalePropNew() *ItemSalePropNew {
	return poolItemSalePropNew.Get().(*ItemSalePropNew)
}

// ReleaseItemSalePropNew 释放ItemSalePropNew
func ReleaseItemSalePropNew(v *ItemSalePropNew) {
	v.SalePropValueStatusList = v.SalePropValueStatusList[:0]
	v.PropertyValue = ""
	v.PropertyId = 0
	poolItemSalePropNew.Put(v)
}
