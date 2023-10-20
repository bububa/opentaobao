package ascp

import (
	"sync"
)

// Insurance 结构体
type Insurance struct {
	// 保险类型
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 保险金额
	Amount string `json:"amount,omitempty" xml:"amount,omitempty"`
}

var poolInsurance = sync.Pool{
	New: func() any {
		return new(Insurance)
	},
}

// GetInsurance() 从对象池中获取Insurance
func GetInsurance() *Insurance {
	return poolInsurance.Get().(*Insurance)
}

// ReleaseInsurance 释放Insurance
func ReleaseInsurance(v *Insurance) {
	v.Type = ""
	v.Amount = ""
	poolInsurance.Put(v)
}
