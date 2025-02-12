package tbk

import (
	"sync"
)

// MaterialUrlList 结构体
type MaterialUrlList struct {
	// 多件价券信息
	MultipleItemsCouponInfoList []MaterialMultiCouponPromotionInfoDto `json:"multiple_items_coupon_info_list,omitempty" xml:"multiple_items_coupon_info_list>material_multi_coupon_promotion_info_dto,omitempty"`
	// 物料对应错误描述
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 入参物料链接
	InputMaterialUrl string `json:"input_material_url,omitempty" xml:"input_material_url,omitempty"`
	// 转链成功的场景下，需要补充说明的信息
	ExtraInfo string `json:"extra_info,omitempty" xml:"extra_info,omitempty"`
	// 物料对应错误码
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
	// 营销信息
	PromotionInfoDto *MaterialPromotionInfoDto `json:"promotion_info_dto,omitempty" xml:"promotion_info_dto,omitempty"`
	// 券信息
	CouponInfoDto *CouponInfoDto `json:"coupon_info_dto,omitempty" xml:"coupon_info_dto,omitempty"`
	// 链接信息
	LinkInfoDto *LinkInfoDto `json:"link_info_dto,omitempty" xml:"link_info_dto,omitempty"`
}

var poolMaterialUrlList = sync.Pool{
	New: func() any {
		return new(MaterialUrlList)
	},
}

// GetMaterialUrlList() 从对象池中获取MaterialUrlList
func GetMaterialUrlList() *MaterialUrlList {
	return poolMaterialUrlList.Get().(*MaterialUrlList)
}

// ReleaseMaterialUrlList 释放MaterialUrlList
func ReleaseMaterialUrlList(v *MaterialUrlList) {
	v.MultipleItemsCouponInfoList = v.MultipleItemsCouponInfoList[:0]
	v.Msg = ""
	v.InputMaterialUrl = ""
	v.ExtraInfo = ""
	v.Code = 0
	v.PromotionInfoDto = nil
	v.CouponInfoDto = nil
	v.LinkInfoDto = nil
	poolMaterialUrlList.Put(v)
}
