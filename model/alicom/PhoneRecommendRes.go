package alicom

import (
	"sync"
)

// PhoneRecommendRes 结构体
type PhoneRecommendRes struct {
	// 商品列表
	List []PhoneDistributionPhoneItemVo `json:"list,omitempty" xml:"list>phone_distribution_phone_item_vo,omitempty"`
	// 公告
	Notice string `json:"notice,omitempty" xml:"notice,omitempty"`
	// 号码信息
	CatInfo *CatInfoVo `json:"cat_info,omitempty" xml:"cat_info,omitempty"`
}

var poolPhoneRecommendRes = sync.Pool{
	New: func() any {
		return new(PhoneRecommendRes)
	},
}

// GetPhoneRecommendRes() 从对象池中获取PhoneRecommendRes
func GetPhoneRecommendRes() *PhoneRecommendRes {
	return poolPhoneRecommendRes.Get().(*PhoneRecommendRes)
}

// ReleasePhoneRecommendRes 释放PhoneRecommendRes
func ReleasePhoneRecommendRes(v *PhoneRecommendRes) {
	v.List = v.List[:0]
	v.Notice = ""
	v.CatInfo = nil
	poolPhoneRecommendRes.Put(v)
}
