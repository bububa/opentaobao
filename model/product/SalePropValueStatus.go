package product

import (
	"sync"
)

// SalePropValueStatus 结构体
type SalePropValueStatus struct {
	// 属性文本
	Text string `json:"text,omitempty" xml:"text,omitempty"`
	// 属性值ID
	Value int64 `json:"value,omitempty" xml:"value,omitempty"`
	// 是否新品
	IsNew bool `json:"is_new,omitempty" xml:"is_new,omitempty"`
}

var poolSalePropValueStatus = sync.Pool{
	New: func() any {
		return new(SalePropValueStatus)
	},
}

// GetSalePropValueStatus() 从对象池中获取SalePropValueStatus
func GetSalePropValueStatus() *SalePropValueStatus {
	return poolSalePropValueStatus.Get().(*SalePropValueStatus)
}

// ReleaseSalePropValueStatus 释放SalePropValueStatus
func ReleaseSalePropValueStatus(v *SalePropValueStatus) {
	v.Text = ""
	v.Value = 0
	v.IsNew = false
	poolSalePropValueStatus.Put(v)
}
