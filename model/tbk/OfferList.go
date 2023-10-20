package tbk

import (
	"sync"
)

// OfferList 结构体
type OfferList struct {
	// 活动id
	OfferId string `json:"offer_id,omitempty" xml:"offer_id,omitempty"`
	// 淘礼金领取URL，不支持使用短链接
	TljUrl string `json:"tlj_url,omitempty" xml:"tlj_url,omitempty"`
	// 商品id
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
}

var poolOfferList = sync.Pool{
	New: func() any {
		return new(OfferList)
	},
}

// GetOfferList() 从对象池中获取OfferList
func GetOfferList() *OfferList {
	return poolOfferList.Get().(*OfferList)
}

// ReleaseOfferList 释放OfferList
func ReleaseOfferList(v *OfferList) {
	v.OfferId = ""
	v.TljUrl = ""
	v.ItemId = ""
	poolOfferList.Put(v)
}
