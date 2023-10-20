package mos

import (
	"sync"
)

// AlibabaMjPresaleSettlementStatisticsData 结构体
type AlibabaMjPresaleSettlementStatisticsData struct {
	// onLineNum
	OnLineNum string `json:"on_line_num,omitempty" xml:"on_line_num,omitempty"`
	// aliPayNum
	AliPayNum string `json:"ali_pay_num,omitempty" xml:"ali_pay_num,omitempty"`
	// offLineNum
	OffLineNum string `json:"off_line_num,omitempty" xml:"off_line_num,omitempty"`
	// aliPayPrice
	AliPayPrice string `json:"ali_pay_price,omitempty" xml:"ali_pay_price,omitempty"`
	// onLinePrice
	OnLinePrice string `json:"on_line_price,omitempty" xml:"on_line_price,omitempty"`
	// offLinePrice
	OffLinePrice string `json:"off_line_price,omitempty" xml:"off_line_price,omitempty"`
}

var poolAlibabaMjPresaleSettlementStatisticsData = sync.Pool{
	New: func() any {
		return new(AlibabaMjPresaleSettlementStatisticsData)
	},
}

// GetAlibabaMjPresaleSettlementStatisticsData() 从对象池中获取AlibabaMjPresaleSettlementStatisticsData
func GetAlibabaMjPresaleSettlementStatisticsData() *AlibabaMjPresaleSettlementStatisticsData {
	return poolAlibabaMjPresaleSettlementStatisticsData.Get().(*AlibabaMjPresaleSettlementStatisticsData)
}

// ReleaseAlibabaMjPresaleSettlementStatisticsData 释放AlibabaMjPresaleSettlementStatisticsData
func ReleaseAlibabaMjPresaleSettlementStatisticsData(v *AlibabaMjPresaleSettlementStatisticsData) {
	v.OnLineNum = ""
	v.AliPayNum = ""
	v.OffLineNum = ""
	v.AliPayPrice = ""
	v.OnLinePrice = ""
	v.OffLinePrice = ""
	poolAlibabaMjPresaleSettlementStatisticsData.Put(v)
}
