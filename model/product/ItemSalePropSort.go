package product

import (
	"sync"
)

// ItemSalePropSort 结构体
type ItemSalePropSort struct {
	// 属性值列表
	SalePropValueSorts []SalePropValueSort `json:"sale_prop_value_sorts,omitempty" xml:"sale_prop_value_sorts>sale_prop_value_sort,omitempty"`
	// 属性项名称
	PropertyValue string `json:"property_value,omitempty" xml:"property_value,omitempty"`
	// 属性项ID
	PropertyId int64 `json:"property_id,omitempty" xml:"property_id,omitempty"`
}

var poolItemSalePropSort = sync.Pool{
	New: func() any {
		return new(ItemSalePropSort)
	},
}

// GetItemSalePropSort() 从对象池中获取ItemSalePropSort
func GetItemSalePropSort() *ItemSalePropSort {
	return poolItemSalePropSort.Get().(*ItemSalePropSort)
}

// ReleaseItemSalePropSort 释放ItemSalePropSort
func ReleaseItemSalePropSort(v *ItemSalePropSort) {
	v.SalePropValueSorts = v.SalePropValueSorts[:0]
	v.PropertyValue = ""
	v.PropertyId = 0
	poolItemSalePropSort.Put(v)
}
