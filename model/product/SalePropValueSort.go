package product

import (
	"sync"
)

// SalePropValueSort 结构体
type SalePropValueSort struct {
	// 属性值文本
	Text string `json:"text,omitempty" xml:"text,omitempty"`
	// 属性值ID
	Value int64 `json:"value,omitempty" xml:"value,omitempty"`
}

var poolSalePropValueSort = sync.Pool{
	New: func() any {
		return new(SalePropValueSort)
	},
}

// GetSalePropValueSort() 从对象池中获取SalePropValueSort
func GetSalePropValueSort() *SalePropValueSort {
	return poolSalePropValueSort.Get().(*SalePropValueSort)
}

// ReleaseSalePropValueSort 释放SalePropValueSort
func ReleaseSalePropValueSort(v *SalePropValueSort) {
	v.Text = ""
	v.Value = 0
	poolSalePropValueSort.Put(v)
}
