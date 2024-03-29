package tbtrade

import (
	"sync"
)

// SendGoodsDetail 结构体
type SendGoodsDetail struct {
	// 包裹详情
	GoodsDetails []PackageGoodsDetail `json:"goods_details,omitempty" xml:"goods_details>package_goods_detail,omitempty"`
	// 0普通1组合商品2组套商品3后置赠品4状态推进
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 0全部发货1部分发货
	ConsignStatus int64 `json:"consign_status,omitempty" xml:"consign_status,omitempty"`
	// 数量
	Amount int64 `json:"amount,omitempty" xml:"amount,omitempty"`
}

var poolSendGoodsDetail = sync.Pool{
	New: func() any {
		return new(SendGoodsDetail)
	},
}

// GetSendGoodsDetail() 从对象池中获取SendGoodsDetail
func GetSendGoodsDetail() *SendGoodsDetail {
	return poolSendGoodsDetail.Get().(*SendGoodsDetail)
}

// ReleaseSendGoodsDetail 释放SendGoodsDetail
func ReleaseSendGoodsDetail(v *SendGoodsDetail) {
	v.GoodsDetails = v.GoodsDetails[:0]
	v.Type = 0
	v.ConsignStatus = 0
	v.Amount = 0
	poolSendGoodsDetail.Put(v)
}
