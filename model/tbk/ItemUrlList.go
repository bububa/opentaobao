package tbk

import (
	"sync"
)

// ItemUrlList 结构体
type ItemUrlList struct {
	// 多件价券信息
	MultipleItemsCouponInfoList []ItemMultiCouponPromotionInfoDto `json:"multiple_items_coupon_info_list,omitempty" xml:"multiple_items_coupon_info_list>item_multi_coupon_promotion_info_dto,omitempty"`
	// 物料对应错误描述
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 入参的商品ID
	InputItemId string `json:"input_item_id,omitempty" xml:"input_item_id,omitempty"`
	// 转链成功的场景下，需要补充说明的信息
	ExtraInfo string `json:"extra_info,omitempty" xml:"extra_info,omitempty"`
	// 物料对应错误码
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
	// 营销信息
	PromotionInfoDto *ItemPromotionInfoDto `json:"promotion_info_dto,omitempty" xml:"promotion_info_dto,omitempty"`
	// 券信息
	CouponInfoDto *CouponInfoDto `json:"coupon_info_dto,omitempty" xml:"coupon_info_dto,omitempty"`
	// 链接信息
	LinkInfoDto *ItemLinkInfoDto `json:"link_info_dto,omitempty" xml:"link_info_dto,omitempty"`
}

var poolItemUrlList = sync.Pool{
	New: func() any {
		return new(ItemUrlList)
	},
}

// GetItemUrlList() 从对象池中获取ItemUrlList
func GetItemUrlList() *ItemUrlList {
	return poolItemUrlList.Get().(*ItemUrlList)
}

// ReleaseItemUrlList 释放ItemUrlList
func ReleaseItemUrlList(v *ItemUrlList) {
	v.MultipleItemsCouponInfoList = v.MultipleItemsCouponInfoList[:0]
	v.Msg = ""
	v.InputItemId = ""
	v.ExtraInfo = ""
	v.Code = 0
	v.PromotionInfoDto = nil
	v.CouponInfoDto = nil
	v.LinkInfoDto = nil
	poolItemUrlList.Put(v)
}
