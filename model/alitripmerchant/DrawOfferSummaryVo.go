package alitripmerchant

import (
	"sync"
)

// DrawOfferSummaryVo 结构体
type DrawOfferSummaryVo struct {
	// 无线端图片
	OfferImageWireless []string `json:"offer_image_wireless,omitempty" xml:"offer_image_wireless>string,omitempty"`
	// 活动名称
	OfferName string `json:"offer_name,omitempty" xml:"offer_name,omitempty"`
	// 子标题
	Subtitle string `json:"subtitle,omitempty" xml:"subtitle,omitempty"`
	// 用户状态
	UserStatus int64 `json:"user_status,omitempty" xml:"user_status,omitempty"`
	// 活动id
	OfferId int64 `json:"offer_id,omitempty" xml:"offer_id,omitempty"`
	// 活动状态
	OfferStatus int64 `json:"offer_status,omitempty" xml:"offer_status,omitempty"`
}

var poolDrawOfferSummaryVo = sync.Pool{
	New: func() any {
		return new(DrawOfferSummaryVo)
	},
}

// GetDrawOfferSummaryVo() 从对象池中获取DrawOfferSummaryVo
func GetDrawOfferSummaryVo() *DrawOfferSummaryVo {
	return poolDrawOfferSummaryVo.Get().(*DrawOfferSummaryVo)
}

// ReleaseDrawOfferSummaryVo 释放DrawOfferSummaryVo
func ReleaseDrawOfferSummaryVo(v *DrawOfferSummaryVo) {
	v.OfferImageWireless = v.OfferImageWireless[:0]
	v.OfferName = ""
	v.Subtitle = ""
	v.UserStatus = 0
	v.OfferId = 0
	v.OfferStatus = 0
	poolDrawOfferSummaryVo.Put(v)
}
