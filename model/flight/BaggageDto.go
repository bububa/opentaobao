package flight

import (
	"sync"
)

// BaggageDto 结构体
type BaggageDto struct {
	// 婴儿车描述
	BabyCar string `json:"baby_car,omitempty" xml:"baby_car,omitempty"`
	// 行李尺寸
	BaggageSize string `json:"baggage_size,omitempty" xml:"baggage_size,omitempty"`
	// 手提行李尺寸
	CarryBagSize string `json:"carry_bag_size,omitempty" xml:"carry_bag_size,omitempty"`
	// 超规行李说明
	ExcessInstruction string `json:"excess_instruction,omitempty" xml:"excess_instruction,omitempty"`
	// 重量单位
	WeightUnit string `json:"weight_unit,omitempty" xml:"weight_unit,omitempty"`
	// 航段
	DepArrDesc string `json:"dep_arr_desc,omitempty" xml:"dep_arr_desc,omitempty"`
	// 乘机人类型
	PassengerTypeDesc string `json:"passenger_type_desc,omitempty" xml:"passenger_type_desc,omitempty"`
	// 行李类型
	BaggageTypeDesc string `json:"baggage_type_desc,omitempty" xml:"baggage_type_desc,omitempty"`
	// 规则描述
	RuleDesc string `json:"rule_desc,omitempty" xml:"rule_desc,omitempty"`
	// 免费行李件数
	BaggageFreePcs int64 `json:"baggage_free_pcs,omitempty" xml:"baggage_free_pcs,omitempty"`
	// 行李重量
	BaggageWeight int64 `json:"baggage_weight,omitempty" xml:"baggage_weight,omitempty"`
	// 手提行李件数
	CarryBagFreePcs int64 `json:"carry_bag_free_pcs,omitempty" xml:"carry_bag_free_pcs,omitempty"`
	// 手提行李重量
	CarryBagWeight int64 `json:"carry_bag_weight,omitempty" xml:"carry_bag_weight,omitempty"`
	// 航程下标 1开始
	OdIndex int64 `json:"od_index,omitempty" xml:"od_index,omitempty"`
	// 乘机人类型，1成人2儿童3婴儿
	PassengerType int64 `json:"passenger_type,omitempty" xml:"passenger_type,omitempty"`
	// 航段下标 1开始
	SegmentIndex int64 `json:"segment_index,omitempty" xml:"segment_index,omitempty"`
	// 手提和托运行李总件数，国内专用
	TotalPcs int64 `json:"total_pcs,omitempty" xml:"total_pcs,omitempty"`
	// 手提和托运行李总重量，国内专用
	TotalWeight int64 `json:"total_weight,omitempty" xml:"total_weight,omitempty"`
	// 托运行李重量是否表示总重量(所有件总量)
	BaggageAllWeight bool `json:"baggage_all_weight,omitempty" xml:"baggage_all_weight,omitempty"`
	// 手提行李是否总重量
	CarryBagAllWeight bool `json:"carry_bag_all_weight,omitempty" xml:"carry_bag_all_weight,omitempty"`
}

var poolBaggageDto = sync.Pool{
	New: func() any {
		return new(BaggageDto)
	},
}

// GetBaggageDto() 从对象池中获取BaggageDto
func GetBaggageDto() *BaggageDto {
	return poolBaggageDto.Get().(*BaggageDto)
}

// ReleaseBaggageDto 释放BaggageDto
func ReleaseBaggageDto(v *BaggageDto) {
	v.BabyCar = ""
	v.BaggageSize = ""
	v.CarryBagSize = ""
	v.ExcessInstruction = ""
	v.WeightUnit = ""
	v.DepArrDesc = ""
	v.PassengerTypeDesc = ""
	v.BaggageTypeDesc = ""
	v.RuleDesc = ""
	v.BaggageFreePcs = 0
	v.BaggageWeight = 0
	v.CarryBagFreePcs = 0
	v.CarryBagWeight = 0
	v.OdIndex = 0
	v.PassengerType = 0
	v.SegmentIndex = 0
	v.TotalPcs = 0
	v.TotalWeight = 0
	v.BaggageAllWeight = false
	v.CarryBagAllWeight = false
	poolBaggageDto.Put(v)
}
