package mos

import (
	"sync"
)

// ResultDo 结构体
type ResultDo struct {
	// data
	Data *CallDispatcherResponse `json:"data,omitempty" xml:"data,omitempty"`
}

var poolResultDo = sync.Pool{
	New: func() any {
		return new(ResultDo)
	},
}

// GetResultDo() 从对象池中获取ResultDo
func GetResultDo() *ResultDo {
	return poolResultDo.Get().(*ResultDo)
}

// ReleaseResultDo 释放ResultDo
func ReleaseResultDo(v *ResultDo) {
	v.Data = nil
	poolResultDo.Put(v)
}
