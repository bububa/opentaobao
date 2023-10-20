package happytrip

import (
	"sync"
)

// ReassignInfo 结构体
type ReassignInfo struct {
	// 改派前订单id，即此订单由哪个订单id改派而生成（当值为空时，表示该订单不是因为改派而生成的）
	PreOrderId string `json:"pre_order_id,omitempty" xml:"pre_order_id,omitempty"`
	// 改派后订单id，即由于订单改派而产生的订单id（当值为空时，表示当前订单未被改派而产生新订单）
	NextOrderId string `json:"next_order_id,omitempty" xml:"next_order_id,omitempty"`
	// 第一个被改派的订单id
	InitOrderId string `json:"init_order_id,omitempty" xml:"init_order_id,omitempty"`
	// 最新被指派的订单id
	LatestOrderId string `json:"latest_order_id,omitempty" xml:"latest_order_id,omitempty"`
}

var poolReassignInfo = sync.Pool{
	New: func() any {
		return new(ReassignInfo)
	},
}

// GetReassignInfo() 从对象池中获取ReassignInfo
func GetReassignInfo() *ReassignInfo {
	return poolReassignInfo.Get().(*ReassignInfo)
}

// ReleaseReassignInfo 释放ReassignInfo
func ReleaseReassignInfo(v *ReassignInfo) {
	v.PreOrderId = ""
	v.NextOrderId = ""
	v.InitOrderId = ""
	v.LatestOrderId = ""
	poolReassignInfo.Put(v)
}
