package dt

import (
	"sync"
)

// RecongnizeItemInfo 结构体
type RecongnizeItemInfo struct {
	// 原产地
	OriginPlace string `json:"origin_place,omitempty" xml:"origin_place,omitempty"`
	// 扩展信息
	ExtMap string `json:"ext_map,omitempty" xml:"ext_map,omitempty"`
	// 优惠结束时间
	PromotionEndTime string `json:"promotion_end_time,omitempty" xml:"promotion_end_time,omitempty"`
	// 优惠开始时间
	PromotionStartTime string `json:"promotion_start_time,omitempty" xml:"promotion_start_time,omitempty"`
	// 规格描述
	SpecDesc string `json:"spec_desc,omitempty" xml:"spec_desc,omitempty"`
	// rt货号
	RtItemNo string `json:"rt_item_no,omitempty" xml:"rt_item_no,omitempty"`
	// 优惠描述
	PromotionDesc string `json:"promotion_desc,omitempty" xml:"promotion_desc,omitempty"`
	// 品牌
	BrandName string `json:"brand_name,omitempty" xml:"brand_name,omitempty"`
	// 品名
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 条码
	BarCode string `json:"bar_code,omitempty" xml:"bar_code,omitempty"`
	// 流水号
	SerialNo string `json:"serial_no,omitempty" xml:"serial_no,omitempty"`
	// 促销类型
	PromotionType int64 `json:"promotion_type,omitempty" xml:"promotion_type,omitempty"`
	// 促销价
	PromotionPrice int64 `json:"promotion_price,omitempty" xml:"promotion_price,omitempty"`
	// 销售价
	SalePrice int64 `json:"sale_price,omitempty" xml:"sale_price,omitempty"`
	// 是否促销中
	InPromotion bool `json:"in_promotion,omitempty" xml:"in_promotion,omitempty"`
}

var poolRecongnizeItemInfo = sync.Pool{
	New: func() any {
		return new(RecongnizeItemInfo)
	},
}

// GetRecongnizeItemInfo() 从对象池中获取RecongnizeItemInfo
func GetRecongnizeItemInfo() *RecongnizeItemInfo {
	return poolRecongnizeItemInfo.Get().(*RecongnizeItemInfo)
}

// ReleaseRecongnizeItemInfo 释放RecongnizeItemInfo
func ReleaseRecongnizeItemInfo(v *RecongnizeItemInfo) {
	v.OriginPlace = ""
	v.ExtMap = ""
	v.PromotionEndTime = ""
	v.PromotionStartTime = ""
	v.SpecDesc = ""
	v.RtItemNo = ""
	v.PromotionDesc = ""
	v.BrandName = ""
	v.Title = ""
	v.BarCode = ""
	v.SerialNo = ""
	v.PromotionType = 0
	v.PromotionPrice = 0
	v.SalePrice = 0
	v.InPromotion = false
	poolRecongnizeItemInfo.Put(v)
}
