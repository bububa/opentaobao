package wdk

import (
	"sync"
)

// PromotionActivityDo 结构体
type PromotionActivityDo struct {
	// 店群机构编码
	SupplyGroupCodes []string `json:"supply_group_codes,omitempty" xml:"supply_group_codes>string,omitempty"`
	// 档期活动开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 创建时间，可不传
	CreateTime string `json:"create_time,omitempty" xml:"create_time,omitempty"`
	// 档期活动名称
	PromotionName string `json:"promotion_name,omitempty" xml:"promotion_name,omitempty"`
	// 档期活动描述
	PromotionDesc string `json:"promotion_desc,omitempty" xml:"promotion_desc,omitempty"`
	// 本期不作区分，可不传
	SupplierCode string `json:"supplier_code,omitempty" xml:"supplier_code,omitempty"`
	// 档期活动结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 本期不作区分，可不传
	OuterPromotionCode string `json:"outer_promotion_code,omitempty" xml:"outer_promotion_code,omitempty"`
	// 盒马帮商家编码
	MerchantCode string `json:"merchant_code,omitempty" xml:"merchant_code,omitempty"`
	// 创建人名称
	Creator string `json:"creator,omitempty" xml:"creator,omitempty"`
	// 创建人ID
	CreatorId int64 `json:"creator_id,omitempty" xml:"creator_id,omitempty"`
}

var poolPromotionActivityDo = sync.Pool{
	New: func() any {
		return new(PromotionActivityDo)
	},
}

// GetPromotionActivityDo() 从对象池中获取PromotionActivityDo
func GetPromotionActivityDo() *PromotionActivityDo {
	return poolPromotionActivityDo.Get().(*PromotionActivityDo)
}

// ReleasePromotionActivityDo 释放PromotionActivityDo
func ReleasePromotionActivityDo(v *PromotionActivityDo) {
	v.SupplyGroupCodes = v.SupplyGroupCodes[:0]
	v.StartTime = ""
	v.CreateTime = ""
	v.PromotionName = ""
	v.PromotionDesc = ""
	v.SupplierCode = ""
	v.EndTime = ""
	v.OuterPromotionCode = ""
	v.MerchantCode = ""
	v.Creator = ""
	v.CreatorId = 0
	poolPromotionActivityDo.Put(v)
}
