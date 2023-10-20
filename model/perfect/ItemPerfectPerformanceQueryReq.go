package perfect

import (
	"sync"
)

// ItemPerfectPerformanceQueryReq 结构体
type ItemPerfectPerformanceQueryReq struct {
	// 天猫商品id
	ItemTmallId string `json:"item_tmall_id,omitempty" xml:"item_tmall_id,omitempty"`
}

var poolItemPerfectPerformanceQueryReq = sync.Pool{
	New: func() any {
		return new(ItemPerfectPerformanceQueryReq)
	},
}

// GetItemPerfectPerformanceQueryReq() 从对象池中获取ItemPerfectPerformanceQueryReq
func GetItemPerfectPerformanceQueryReq() *ItemPerfectPerformanceQueryReq {
	return poolItemPerfectPerformanceQueryReq.Get().(*ItemPerfectPerformanceQueryReq)
}

// ReleaseItemPerfectPerformanceQueryReq 释放ItemPerfectPerformanceQueryReq
func ReleaseItemPerfectPerformanceQueryReq(v *ItemPerfectPerformanceQueryReq) {
	v.ItemTmallId = ""
	poolItemPerfectPerformanceQueryReq.Put(v)
}
