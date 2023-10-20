package xhotelonlineorder

import (
	"sync"
)

// OtherFeeInfo 结构体
type OtherFeeInfo struct {
	// 杂费名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 杂费金额
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
	// 结账状态(1:结账,0:未结账)
	Checkout int64 `json:"checkout,omitempty" xml:"checkout,omitempty"`
}

var poolOtherFeeInfo = sync.Pool{
	New: func() any {
		return new(OtherFeeInfo)
	},
}

// GetOtherFeeInfo() 从对象池中获取OtherFeeInfo
func GetOtherFeeInfo() *OtherFeeInfo {
	return poolOtherFeeInfo.Get().(*OtherFeeInfo)
}

// ReleaseOtherFeeInfo 释放OtherFeeInfo
func ReleaseOtherFeeInfo(v *OtherFeeInfo) {
	v.Name = ""
	v.Price = 0
	v.Checkout = 0
	poolOtherFeeInfo.Put(v)
}
