package ieagency

import (
	"sync"
)

// IeItemVo 结构体
type IeItemVo struct {
	// 航班信息
	Flights []IeFlightVo `json:"flights,omitempty" xml:"flights>ie_flight_vo,omitempty"`
	// 行李规定
	BaggageRule string `json:"baggage_rule,omitempty" xml:"baggage_rule,omitempty"`
	// bonusId
	BonusId string `json:"bonus_id,omitempty" xml:"bonus_id,omitempty"`
	// fareItemId
	FareItemId string `json:"fare_item_id,omitempty" xml:"fare_item_id,omitempty"`
	// officeNo
	OfficeNo string `json:"office_no,omitempty" xml:"office_no,omitempty"`
	// 供应渠道
	ResourceCode string `json:"resource_code,omitempty" xml:"resource_code,omitempty"`
	// 退改签规则
	Restriction string `json:"restriction,omitempty" xml:"restriction,omitempty"`
	// 出票航司
	TicketingAirline string `json:"ticketing_airline,omitempty" xml:"ticketing_airline,omitempty"`
	// 行程类型(OneWay:单程,RoundTrip:往返,MultiCity:多程)
	TripType string `json:"trip_type,omitempty" xml:"trip_type,omitempty"`
	// 原始政策id
	OriginBonusId string `json:"origin_bonus_id,omitempty" xml:"origin_bonus_id,omitempty"`
	// 成人价格
	AdultPrice int64 `json:"adult_price,omitempty" xml:"adult_price,omitempty"`
	// 成人价格
	AdultTax int64 `json:"adult_tax,omitempty" xml:"adult_tax,omitempty"`
	// 儿童价格
	ChildPrice int64 `json:"child_price,omitempty" xml:"child_price,omitempty"`
	// 儿童税费
	ChildTax int64 `json:"child_tax,omitempty" xml:"child_tax,omitempty"`
	// 婴儿税费
	InfantTax int64 `json:"infant_tax,omitempty" xml:"infant_tax,omitempty"`
	// 婴儿价格
	InfantPrice int64 `json:"infant_price,omitempty" xml:"infant_price,omitempty"`
}

var poolIeItemVo = sync.Pool{
	New: func() any {
		return new(IeItemVo)
	},
}

// GetIeItemVo() 从对象池中获取IeItemVo
func GetIeItemVo() *IeItemVo {
	return poolIeItemVo.Get().(*IeItemVo)
}

// ReleaseIeItemVo 释放IeItemVo
func ReleaseIeItemVo(v *IeItemVo) {
	v.Flights = v.Flights[:0]
	v.BaggageRule = ""
	v.BonusId = ""
	v.FareItemId = ""
	v.OfficeNo = ""
	v.ResourceCode = ""
	v.Restriction = ""
	v.TicketingAirline = ""
	v.TripType = ""
	v.OriginBonusId = ""
	v.AdultPrice = 0
	v.AdultTax = 0
	v.ChildPrice = 0
	v.ChildTax = 0
	v.InfantTax = 0
	v.InfantPrice = 0
	poolIeItemVo.Put(v)
}
