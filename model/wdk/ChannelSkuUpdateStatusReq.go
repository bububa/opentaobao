package wdk

import (
	"sync"
)

// ChannelSkuUpdateStatusReq 结构体
type ChannelSkuUpdateStatusReq struct {
	// 商品编码
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// 经营店编码
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 渠道编码：2 美团，3 饿了么，26 京东到家，31 翱象淘鲜达，32 翱象共享库存
	ChannelCode string `json:"channel_code,omitempty" xml:"channel_code,omitempty"`
	// 是否渠道可售 1可售（上架）0：不可售（下架）
	OnlineSaleFlag int64 `json:"online_sale_flag,omitempty" xml:"online_sale_flag,omitempty"`
}

var poolChannelSkuUpdateStatusReq = sync.Pool{
	New: func() any {
		return new(ChannelSkuUpdateStatusReq)
	},
}

// GetChannelSkuUpdateStatusReq() 从对象池中获取ChannelSkuUpdateStatusReq
func GetChannelSkuUpdateStatusReq() *ChannelSkuUpdateStatusReq {
	return poolChannelSkuUpdateStatusReq.Get().(*ChannelSkuUpdateStatusReq)
}

// ReleaseChannelSkuUpdateStatusReq 释放ChannelSkuUpdateStatusReq
func ReleaseChannelSkuUpdateStatusReq(v *ChannelSkuUpdateStatusReq) {
	v.SkuCode = ""
	v.StoreId = ""
	v.ChannelCode = ""
	v.OnlineSaleFlag = 0
	poolChannelSkuUpdateStatusReq.Put(v)
}
