package ieagency

import (
	"sync"
)

// IeOrderVo 结构体
type IeOrderVo struct {
	// activityVOs
	ActivityVos []IeOrderActivityVo `json:"activity_vos,omitempty" xml:"activity_vos>ie_order_activity_vo,omitempty"`
	// 预定订单信息
	BookOrderVos []IeBookOrderVo `json:"book_order_vos,omitempty" xml:"book_order_vos>ie_book_order_vo,omitempty"`
	// 乘机人列表
	PassgenerVos []IePassgenerVo `json:"passgener_vos,omitempty" xml:"passgener_vos>ie_passgener_vo,omitempty"`
	// 辅营订单
	VirProOrderVos []IeVirProOrderVo `json:"vir_pro_order_vos,omitempty" xml:"vir_pro_order_vos>ie_vir_pro_order_vo,omitempty"`
	// 订单基本信息
	BaseOrderVo *IeBaseOrderVo `json:"base_order_vo,omitempty" xml:"base_order_vo,omitempty"`
	// 联系人信息
	ContactsVo *IeContactsVo `json:"contacts_vo,omitempty" xml:"contacts_vo,omitempty"`
	// 搜索产品信息
	ItemVo *IeItemVo `json:"item_vo,omitempty" xml:"item_vo,omitempty"`
	// 行程单信息
	ItineraryVo *IeItineraryVo `json:"itinerary_vo,omitempty" xml:"itinerary_vo,omitempty"`
	// 交易订单ID
	TradeOrderId int64 `json:"trade_order_id,omitempty" xml:"trade_order_id,omitempty"`
	// 订单模式
	BusinessMode int64 `json:"business_mode,omitempty" xml:"business_mode,omitempty"`
}

var poolIeOrderVo = sync.Pool{
	New: func() any {
		return new(IeOrderVo)
	},
}

// GetIeOrderVo() 从对象池中获取IeOrderVo
func GetIeOrderVo() *IeOrderVo {
	return poolIeOrderVo.Get().(*IeOrderVo)
}

// ReleaseIeOrderVo 释放IeOrderVo
func ReleaseIeOrderVo(v *IeOrderVo) {
	v.ActivityVos = v.ActivityVos[:0]
	v.BookOrderVos = v.BookOrderVos[:0]
	v.PassgenerVos = v.PassgenerVos[:0]
	v.VirProOrderVos = v.VirProOrderVos[:0]
	v.BaseOrderVo = nil
	v.ContactsVo = nil
	v.ItemVo = nil
	v.ItineraryVo = nil
	v.TradeOrderId = 0
	v.BusinessMode = 0
	poolIeOrderVo.Put(v)
}
