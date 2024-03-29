package singletreasure

import (
	"sync"
)

// ItemDetailInfo 结构体
type ItemDetailInfo struct {
	// 活动结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 不包邮地区
	ExcludeAreas string `json:"exclude_areas,omitempty" xml:"exclude_areas,omitempty"`
	// 活动开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 活动名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 对应的人群 id
	CrowdId string `json:"crowd_id,omitempty" xml:"crowd_id,omitempty"`
	// 限购数
	LimitCheck int64 `json:"limit_check,omitempty" xml:"limit_check,omitempty"`
	// 活动 id
	ActivityId int64 `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
	// 折扣类型，1 为促销价，2 为打折，4 为减钱
	DiscountType int64 `json:"discount_type,omitempty" xml:"discount_type,omitempty"`
	// 猫客对应折扣
	MkDiscount int64 `json:"mk_discount,omitempty" xml:"mk_discount,omitempty"`
	// 优惠类型，1 为商品级优惠，3 为 SKU 级优惠
	PromotionLevel int64 `json:"promotion_level,omitempty" xml:"promotion_level,omitempty"`
	// 状态，-1 删除，0 暂停，1 征程
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 宝贝 id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 折扣类型，1 为促销价，2 为打折，4 为减钱
	Discount int64 `json:"discount,omitempty" xml:"discount,omitempty"`
	// 价格
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
	// id
	DetailId int64 `json:"detail_id,omitempty" xml:"detail_id,omitempty"`
	// skuId
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 是否包邮
	FreePost bool `json:"free_post,omitempty" xml:"free_post,omitempty"`
	// 抹分与否
	IsDiscardFen bool `json:"is_discard_fen,omitempty" xml:"is_discard_fen,omitempty"`
	// 是否取整
	IsMathFloor bool `json:"is_math_floor,omitempty" xml:"is_math_floor,omitempty"`
}

var poolItemDetailInfo = sync.Pool{
	New: func() any {
		return new(ItemDetailInfo)
	},
}

// GetItemDetailInfo() 从对象池中获取ItemDetailInfo
func GetItemDetailInfo() *ItemDetailInfo {
	return poolItemDetailInfo.Get().(*ItemDetailInfo)
}

// ReleaseItemDetailInfo 释放ItemDetailInfo
func ReleaseItemDetailInfo(v *ItemDetailInfo) {
	v.EndTime = ""
	v.ExcludeAreas = ""
	v.StartTime = ""
	v.Name = ""
	v.CrowdId = ""
	v.LimitCheck = 0
	v.ActivityId = 0
	v.DiscountType = 0
	v.MkDiscount = 0
	v.PromotionLevel = 0
	v.Status = 0
	v.ItemId = 0
	v.Discount = 0
	v.Price = 0
	v.DetailId = 0
	v.SkuId = 0
	v.FreePost = false
	v.IsDiscardFen = false
	v.IsMathFloor = false
	poolItemDetailInfo.Put(v)
}
