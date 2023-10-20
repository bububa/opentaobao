package alsc

import (
	"sync"
)

// QueryPhyCardOpenReq 结构体
type QueryPhyCardOpenReq struct {
	// 外部品牌ID
	OutBrandId string `json:"out_brand_id,omitempty" xml:"out_brand_id,omitempty"`
	// 物理卡号
	PhysicalCardId string `json:"physical_card_id,omitempty" xml:"physical_card_id,omitempty"`
	// 品牌ID
	BrandId string `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
}

var poolQueryPhyCardOpenReq = sync.Pool{
	New: func() any {
		return new(QueryPhyCardOpenReq)
	},
}

// GetQueryPhyCardOpenReq() 从对象池中获取QueryPhyCardOpenReq
func GetQueryPhyCardOpenReq() *QueryPhyCardOpenReq {
	return poolQueryPhyCardOpenReq.Get().(*QueryPhyCardOpenReq)
}

// ReleaseQueryPhyCardOpenReq 释放QueryPhyCardOpenReq
func ReleaseQueryPhyCardOpenReq(v *QueryPhyCardOpenReq) {
	v.OutBrandId = ""
	v.PhysicalCardId = ""
	v.BrandId = ""
	poolQueryPhyCardOpenReq.Put(v)
}
