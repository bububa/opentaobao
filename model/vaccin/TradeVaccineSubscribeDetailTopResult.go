package vaccin

import (
	"sync"
)

// TradeVaccineSubscribeDetailTopResult 结构体
type TradeVaccineSubscribeDetailTopResult struct {
	// 预约单模型
	DetailTopDtoList []TradeVaccineSubscribeDetailTopDto `json:"detail_top_dto_list,omitempty" xml:"detail_top_dto_list>trade_vaccine_subscribe_detail_top_dto,omitempty"`
}

var poolTradeVaccineSubscribeDetailTopResult = sync.Pool{
	New: func() any {
		return new(TradeVaccineSubscribeDetailTopResult)
	},
}

// GetTradeVaccineSubscribeDetailTopResult() 从对象池中获取TradeVaccineSubscribeDetailTopResult
func GetTradeVaccineSubscribeDetailTopResult() *TradeVaccineSubscribeDetailTopResult {
	return poolTradeVaccineSubscribeDetailTopResult.Get().(*TradeVaccineSubscribeDetailTopResult)
}

// ReleaseTradeVaccineSubscribeDetailTopResult 释放TradeVaccineSubscribeDetailTopResult
func ReleaseTradeVaccineSubscribeDetailTopResult(v *TradeVaccineSubscribeDetailTopResult) {
	v.DetailTopDtoList = v.DetailTopDtoList[:0]
	poolTradeVaccineSubscribeDetailTopResult.Put(v)
}
