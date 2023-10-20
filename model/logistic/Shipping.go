package logistic

import (
	"sync"
)

// Shipping 结构体
type Shipping struct {
	// 返回发货是否成功。
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

var poolShipping = sync.Pool{
	New: func() any {
		return new(Shipping)
	},
}

// GetShipping() 从对象池中获取Shipping
func GetShipping() *Shipping {
	return poolShipping.Get().(*Shipping)
}

// ReleaseShipping 释放Shipping
func ReleaseShipping(v *Shipping) {
	v.IsSuccess = false
	poolShipping.Put(v)
}
