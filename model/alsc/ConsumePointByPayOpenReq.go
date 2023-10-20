package alsc

import (
	"sync"
)

// ConsumePointByPayOpenReq 结构体
type ConsumePointByPayOpenReq struct {
	// 品牌id
	BrandId string `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
	// 顾客id
	CustomerId string `json:"customer_id,omitempty" xml:"customer_id,omitempty"`
	// 外部门店ID
	OutShopId string `json:"out_shop_id,omitempty" xml:"out_shop_id,omitempty"`
	// 外部品牌id,brandId与out_brand_id不可同时为空
	OutBrandId string `json:"out_brand_id,omitempty" xml:"out_brand_id,omitempty"`
	// 门店id
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 变更积分
	ChangePoint int64 `json:"change_point,omitempty" xml:"change_point,omitempty"`
}

var poolConsumePointByPayOpenReq = sync.Pool{
	New: func() any {
		return new(ConsumePointByPayOpenReq)
	},
}

// GetConsumePointByPayOpenReq() 从对象池中获取ConsumePointByPayOpenReq
func GetConsumePointByPayOpenReq() *ConsumePointByPayOpenReq {
	return poolConsumePointByPayOpenReq.Get().(*ConsumePointByPayOpenReq)
}

// ReleaseConsumePointByPayOpenReq 释放ConsumePointByPayOpenReq
func ReleaseConsumePointByPayOpenReq(v *ConsumePointByPayOpenReq) {
	v.BrandId = ""
	v.CustomerId = ""
	v.OutShopId = ""
	v.OutBrandId = ""
	v.ShopId = ""
	v.ChangePoint = 0
	poolConsumePointByPayOpenReq.Put(v)
}
