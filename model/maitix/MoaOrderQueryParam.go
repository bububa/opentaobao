package maitix

import (
	"sync"
)

// MoaOrderQueryParam 结构体
type MoaOrderQueryParam struct {
	// 大麦订单id
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 查询是否排除取消的订单
	ExcludeUselessSubOrder bool `json:"exclude_useless_sub_order,omitempty" xml:"exclude_useless_sub_order,omitempty"`
}

var poolMoaOrderQueryParam = sync.Pool{
	New: func() any {
		return new(MoaOrderQueryParam)
	},
}

// GetMoaOrderQueryParam() 从对象池中获取MoaOrderQueryParam
func GetMoaOrderQueryParam() *MoaOrderQueryParam {
	return poolMoaOrderQueryParam.Get().(*MoaOrderQueryParam)
}

// ReleaseMoaOrderQueryParam 释放MoaOrderQueryParam
func ReleaseMoaOrderQueryParam(v *MoaOrderQueryParam) {
	v.OrderId = 0
	v.ExcludeUselessSubOrder = false
	poolMoaOrderQueryParam.Put(v)
}
