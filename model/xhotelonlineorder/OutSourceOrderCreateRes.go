package xhotelonlineorder

import (
	"sync"
)

// OutSourceOrderCreateRes 结构体
type OutSourceOrderCreateRes struct {
	// 飞猪酒店订单号
	BizOrderId int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
}

var poolOutSourceOrderCreateRes = sync.Pool{
	New: func() any {
		return new(OutSourceOrderCreateRes)
	},
}

// GetOutSourceOrderCreateRes() 从对象池中获取OutSourceOrderCreateRes
func GetOutSourceOrderCreateRes() *OutSourceOrderCreateRes {
	return poolOutSourceOrderCreateRes.Get().(*OutSourceOrderCreateRes)
}

// ReleaseOutSourceOrderCreateRes 释放OutSourceOrderCreateRes
func ReleaseOutSourceOrderCreateRes(v *OutSourceOrderCreateRes) {
	v.BizOrderId = 0
	poolOutSourceOrderCreateRes.Put(v)
}
