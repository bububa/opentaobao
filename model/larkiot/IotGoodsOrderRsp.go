package larkiot

import (
	"sync"
)

// IotGoodsOrderRsp 结构体
type IotGoodsOrderRsp struct {
	// 订单号
	GoodsOrderId string `json:"goods_order_id,omitempty" xml:"goods_order_id,omitempty"`
	// 影院内码
	CinemaLinkId string `json:"cinema_link_id,omitempty" xml:"cinema_link_id,omitempty"`
	// 外部订单号
	OutGoodsOrderId string `json:"out_goods_order_id,omitempty" xml:"out_goods_order_id,omitempty"`
}

var poolIotGoodsOrderRsp = sync.Pool{
	New: func() any {
		return new(IotGoodsOrderRsp)
	},
}

// GetIotGoodsOrderRsp() 从对象池中获取IotGoodsOrderRsp
func GetIotGoodsOrderRsp() *IotGoodsOrderRsp {
	return poolIotGoodsOrderRsp.Get().(*IotGoodsOrderRsp)
}

// ReleaseIotGoodsOrderRsp 释放IotGoodsOrderRsp
func ReleaseIotGoodsOrderRsp(v *IotGoodsOrderRsp) {
	v.GoodsOrderId = ""
	v.CinemaLinkId = ""
	v.OutGoodsOrderId = ""
	poolIotGoodsOrderRsp.Put(v)
}
