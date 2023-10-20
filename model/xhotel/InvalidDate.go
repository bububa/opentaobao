package xhotel

import (
	"sync"
)

// InvalidDate 结构体
type InvalidDate struct {
	// 活动失效开始时间
	InvalidFrom string `json:"invalid_from,omitempty" xml:"invalid_from,omitempty"`
	// 活动失效结束时间
	InvalidTo string `json:"invalid_to,omitempty" xml:"invalid_to,omitempty"`
}

var poolInvalidDate = sync.Pool{
	New: func() any {
		return new(InvalidDate)
	},
}

// GetInvalidDate() 从对象池中获取InvalidDate
func GetInvalidDate() *InvalidDate {
	return poolInvalidDate.Get().(*InvalidDate)
}

// ReleaseInvalidDate 释放InvalidDate
func ReleaseInvalidDate(v *InvalidDate) {
	v.InvalidFrom = ""
	v.InvalidTo = ""
	poolInvalidDate.Put(v)
}
