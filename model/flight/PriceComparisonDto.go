package flight

import (
	"sync"
)

// PriceComparisonDto 结构体
type PriceComparisonDto struct {
	// 航班日期
	TravelDateStr []string `json:"travel_date_str,omitempty" xml:"travel_date_str>string,omitempty"`
	// 政策id
	PolicyIdStr string `json:"policy_id_str,omitempty" xml:"policy_id_str,omitempty"`
	// 出发-到达
	ArrDep string `json:"arr_dep,omitempty" xml:"arr_dep,omitempty"`
	// 航班号
	FlightNos string `json:"flight_nos,omitempty" xml:"flight_nos,omitempty"`
	// 舱位代码
	CarbinList string `json:"carbin_list,omitempty" xml:"carbin_list,omitempty"`
	// 政策类型：0，普通政策；1，特殊政策；2，规则政策
	PolicyType int64 `json:"policy_type,omitempty" xml:"policy_type,omitempty"`
	// 自己投放价格，单位：分
	SelfSalePrice int64 `json:"self_sale_price,omitempty" xml:"self_sale_price,omitempty"`
	// 最优价格，单位：分
	LowestSalePrice int64 `json:"lowest_sale_price,omitempty" xml:"lowest_sale_price,omitempty"`
	// 价差，单位：分
	PriceDiff int64 `json:"price_diff,omitempty" xml:"price_diff,omitempty"`
	// 销售方式：0，无；1，打包销售套餐1；2，打包销售套餐2；3，打包销售套餐3；4，返现-航司运价；5，返现-销售方包装
	SaleModeCode int64 `json:"sale_mode_code,omitempty" xml:"sale_mode_code,omitempty"`
	// 产品类型：0，普通；1，小团；2，学生；3，青年；4，老年；5，地区；6，会员；10，学生认证；11，年龄
	ProductType int64 `json:"product_type,omitempty" xml:"product_type,omitempty"`
	// 政策投放情况：0，未投放；1，已投放，非最优惠；2，已投放，且为最优惠政策
	PolicyDeployStatus int64 `json:"policy_deploy_status,omitempty" xml:"policy_deploy_status,omitempty"`
	// 是否销售
	CanSell bool `json:"can_sell,omitempty" xml:"can_sell,omitempty"`
}

var poolPriceComparisonDto = sync.Pool{
	New: func() any {
		return new(PriceComparisonDto)
	},
}

// GetPriceComparisonDto() 从对象池中获取PriceComparisonDto
func GetPriceComparisonDto() *PriceComparisonDto {
	return poolPriceComparisonDto.Get().(*PriceComparisonDto)
}

// ReleasePriceComparisonDto 释放PriceComparisonDto
func ReleasePriceComparisonDto(v *PriceComparisonDto) {
	v.TravelDateStr = v.TravelDateStr[:0]
	v.PolicyIdStr = ""
	v.ArrDep = ""
	v.FlightNos = ""
	v.CarbinList = ""
	v.PolicyType = 0
	v.SelfSalePrice = 0
	v.LowestSalePrice = 0
	v.PriceDiff = 0
	v.SaleModeCode = 0
	v.ProductType = 0
	v.PolicyDeployStatus = 0
	v.CanSell = false
	poolPriceComparisonDto.Put(v)
}
