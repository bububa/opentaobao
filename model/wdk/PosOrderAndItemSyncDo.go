package wdk

import (
	"sync"
)

// PosOrderAndItemSyncDo 结构体
type PosOrderAndItemSyncDo struct {
	// 商品信息
	ItemInfos []ItemInfo `json:"item_infos,omitempty" xml:"item_infos>item_info,omitempty"`
	// 订单流水信息
	OrderInfo *OrderInfoDo `json:"order_info,omitempty" xml:"order_info,omitempty"`
}

var poolPosOrderAndItemSyncDo = sync.Pool{
	New: func() any {
		return new(PosOrderAndItemSyncDo)
	},
}

// GetPosOrderAndItemSyncDo() 从对象池中获取PosOrderAndItemSyncDo
func GetPosOrderAndItemSyncDo() *PosOrderAndItemSyncDo {
	return poolPosOrderAndItemSyncDo.Get().(*PosOrderAndItemSyncDo)
}

// ReleasePosOrderAndItemSyncDo 释放PosOrderAndItemSyncDo
func ReleasePosOrderAndItemSyncDo(v *PosOrderAndItemSyncDo) {
	v.ItemInfos = v.ItemInfos[:0]
	v.OrderInfo = nil
	poolPosOrderAndItemSyncDo.Put(v)
}
