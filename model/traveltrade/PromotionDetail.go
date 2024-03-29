package traveltrade

import (
	"sync"
)

// PromotionDetail 结构体
type PromotionDetail struct {
	// 优惠金额（免运费、限时打折时为空）,单位：分
	DiscountFee string `json:"discount_fee,omitempty" xml:"discount_fee,omitempty"`
	// 优惠活动的描述
	PromotionDesc string `json:"promotion_desc,omitempty" xml:"promotion_desc,omitempty"`
	// 优惠id，(由营销工具id、优惠活动id和优惠详情id组成，结构为：营销工具id-优惠活动id_优惠详情id，如mjs-123024_211143）
	PromotionId string `json:"promotion_id,omitempty" xml:"promotion_id,omitempty"`
	// 优惠信息的名称
	PromotionName string `json:"promotion_name,omitempty" xml:"promotion_name,omitempty"`
	// 交易的主订单或子订单号
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolPromotionDetail = sync.Pool{
	New: func() any {
		return new(PromotionDetail)
	},
}

// GetPromotionDetail() 从对象池中获取PromotionDetail
func GetPromotionDetail() *PromotionDetail {
	return poolPromotionDetail.Get().(*PromotionDetail)
}

// ReleasePromotionDetail 释放PromotionDetail
func ReleasePromotionDetail(v *PromotionDetail) {
	v.DiscountFee = ""
	v.PromotionDesc = ""
	v.PromotionId = ""
	v.PromotionName = ""
	v.Id = 0
	poolPromotionDetail.Put(v)
}
