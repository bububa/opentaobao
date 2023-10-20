package bus

import (
	"sync"
)

// InsuranceRecommendRq 结构体
type InsuranceRecommendRq struct {
	// 出发城市
	CityName string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	// 机具编号
	MachineNumber string `json:"machine_number,omitempty" xml:"machine_number,omitempty"`
	// 出发站点
	StationName string `json:"station_name,omitempty" xml:"station_name,omitempty"`
	// 出发省份
	ProvinceName string `json:"province_name,omitempty" xml:"province_name,omitempty"`
	// 渠道来源：self-&gt;自助机机具;window-&gt;窗口
	TradeSource string `json:"trade_source,omitempty" xml:"trade_source,omitempty"`
}

var poolInsuranceRecommendRq = sync.Pool{
	New: func() any {
		return new(InsuranceRecommendRq)
	},
}

// GetInsuranceRecommendRq() 从对象池中获取InsuranceRecommendRq
func GetInsuranceRecommendRq() *InsuranceRecommendRq {
	return poolInsuranceRecommendRq.Get().(*InsuranceRecommendRq)
}

// ReleaseInsuranceRecommendRq 释放InsuranceRecommendRq
func ReleaseInsuranceRecommendRq(v *InsuranceRecommendRq) {
	v.CityName = ""
	v.MachineNumber = ""
	v.StationName = ""
	v.ProvinceName = ""
	v.TradeSource = ""
	poolInsuranceRecommendRq.Put(v)
}
