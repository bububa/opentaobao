package maitix

import (
	"sync"
)

// EticketQueryParam 结构体
type EticketQueryParam struct {
	// 主分销单Id 必传
	MainOrderId int64 `json:"main_order_id,omitempty" xml:"main_order_id,omitempty"`
	// 子分销单Id 可传可不传
	SubOrderId int64 `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
}

var poolEticketQueryParam = sync.Pool{
	New: func() any {
		return new(EticketQueryParam)
	},
}

// GetEticketQueryParam() 从对象池中获取EticketQueryParam
func GetEticketQueryParam() *EticketQueryParam {
	return poolEticketQueryParam.Get().(*EticketQueryParam)
}

// ReleaseEticketQueryParam 释放EticketQueryParam
func ReleaseEticketQueryParam(v *EticketQueryParam) {
	v.MainOrderId = 0
	v.SubOrderId = 0
	poolEticketQueryParam.Put(v)
}
