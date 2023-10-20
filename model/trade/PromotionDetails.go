package trade

import (
	"sync"
)

// PromotionDetails 结构体
type PromotionDetails struct {
	// 优惠信息的名称
	PromotionName string `json:"promotion_name,omitempty" xml:"promotion_name,omitempty"`
	// 优惠金额（免运费、限时打折时为空）,单位：元
	DiscountFee string `json:"discount_fee,omitempty" xml:"discount_fee,omitempty"`
	// 满就送商品时，所送商品的名称
	GiftItemName string `json:"gift_item_name,omitempty" xml:"gift_item_name,omitempty"`
	// 赠品的宝贝id
	GiftItemId string `json:"gift_item_id,omitempty" xml:"gift_item_id,omitempty"`
	// 满就送礼物的礼物数量
	GiftItemNum string `json:"gift_item_num,omitempty" xml:"gift_item_num,omitempty"`
	// 优惠活动的描述
	PromotionDesc string `json:"promotion_desc,omitempty" xml:"promotion_desc,omitempty"`
	// 优惠id，(由营销工具id、优惠活动id和优惠详情id组成，结构为：营销工具id-优惠活动id_优惠详情id，如mjs-123024_211143）
	PromotionId string `json:"promotion_id,omitempty" xml:"promotion_id,omitempty"`
	// 交易的主订单或子订单号
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolPromotionDetails = sync.Pool{
	New: func() any {
		return new(PromotionDetails)
	},
}

// GetPromotionDetails() 从对象池中获取PromotionDetails
func GetPromotionDetails() *PromotionDetails {
	return poolPromotionDetails.Get().(*PromotionDetails)
}

// ReleasePromotionDetails 释放PromotionDetails
func ReleasePromotionDetails(v *PromotionDetails) {
	v.PromotionName = ""
	v.DiscountFee = ""
	v.GiftItemName = ""
	v.GiftItemId = ""
	v.GiftItemNum = ""
	v.PromotionDesc = ""
	v.PromotionId = ""
	v.Id = 0
	poolPromotionDetails.Put(v)
}
