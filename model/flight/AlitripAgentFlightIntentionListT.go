package flight

import (
	"sync"
)

// AlitripAgentFlightIntentionListT 结构体
type AlitripAgentFlightIntentionListT struct {
	// 乘机人信息
	PassengerItemInfos []PassengerInfosDto `json:"passenger_item_infos,omitempty" xml:"passenger_item_infos>passenger_infos_dto,omitempty"`
	// 行李
	BaggageRule *BaggageRuleDto `json:"baggage_rule,omitempty" xml:"baggage_rule,omitempty"`
	// 国内国际
	DomesticIntl int64 `json:"domestic_intl,omitempty" xml:"domestic_intl,omitempty"`
	// 航班信息
	FlightInfo *FlightInfoDto `json:"flight_info,omitempty" xml:"flight_info,omitempty"`
	// 主键
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 退改签规则
	PenaltyRule *PenaltyRuleDto `json:"penalty_rule,omitempty" xml:"penalty_rule,omitempty"`
	// 总机建费(分)
	TotalBuildPrice int64 `json:"total_build_price,omitempty" xml:"total_build_price,omitempty"`
	// 总燃油费(分)
	TotalOilPrice int64 `json:"total_oil_price,omitempty" xml:"total_oil_price,omitempty"`
	// 总价格(分)
	TotalPrice int64 `json:"total_price,omitempty" xml:"total_price,omitempty"`
}

var poolAlitripAgentFlightIntentionListT = sync.Pool{
	New: func() any {
		return new(AlitripAgentFlightIntentionListT)
	},
}

// GetAlitripAgentFlightIntentionListT() 从对象池中获取AlitripAgentFlightIntentionListT
func GetAlitripAgentFlightIntentionListT() *AlitripAgentFlightIntentionListT {
	return poolAlitripAgentFlightIntentionListT.Get().(*AlitripAgentFlightIntentionListT)
}

// ReleaseAlitripAgentFlightIntentionListT 释放AlitripAgentFlightIntentionListT
func ReleaseAlitripAgentFlightIntentionListT(v *AlitripAgentFlightIntentionListT) {
	v.PassengerItemInfos = v.PassengerItemInfos[:0]
	v.BaggageRule = nil
	v.DomesticIntl = 0
	v.FlightInfo = nil
	v.Id = 0
	v.PenaltyRule = nil
	v.TotalBuildPrice = 0
	v.TotalOilPrice = 0
	v.TotalPrice = 0
	poolAlitripAgentFlightIntentionListT.Put(v)
}
