package ieagency

import (
	"sync"
)

// ItemParam 结构体
type ItemParam struct {
	// 航班行程信息
	ItineraryParams []ItineraryParam `json:"itinerary_params,omitempty" xml:"itinerary_params>itinerary_param,omitempty"`
	// 成人税费（单位：分）
	AdultTax int64 `json:"adult_tax,omitempty" xml:"adult_tax,omitempty"`
	// 成人票价（单位：分）
	AdultTicketPrice int64 `json:"adult_ticket_price,omitempty" xml:"adult_ticket_price,omitempty"`
	// 儿童税费（单位：分）
	ChildTax int64 `json:"child_tax,omitempty" xml:"child_tax,omitempty"`
	// 儿童票价（单位：分）
	ChildTicketPrice int64 `json:"child_ticket_price,omitempty" xml:"child_ticket_price,omitempty"`
	// 行程类型(1:单程;2:往返; 3:多程）
	TripType int64 `json:"trip_type,omitempty" xml:"trip_type,omitempty"`
}

var poolItemParam = sync.Pool{
	New: func() any {
		return new(ItemParam)
	},
}

// GetItemParam() 从对象池中获取ItemParam
func GetItemParam() *ItemParam {
	return poolItemParam.Get().(*ItemParam)
}

// ReleaseItemParam 释放ItemParam
func ReleaseItemParam(v *ItemParam) {
	v.ItineraryParams = v.ItineraryParams[:0]
	v.AdultTax = 0
	v.AdultTicketPrice = 0
	v.ChildTax = 0
	v.ChildTicketPrice = 0
	v.TripType = 0
	poolItemParam.Put(v)
}
