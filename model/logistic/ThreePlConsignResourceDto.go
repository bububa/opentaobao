package logistic

import (
	"sync"
)

// ThreePlConsignResourceDto 结构体
type ThreePlConsignResourceDto struct {
	// 资源code
	ResCode string `json:"res_code,omitempty" xml:"res_code,omitempty"`
	// 资源名称
	ResName string `json:"res_name,omitempty" xml:"res_name,omitempty"`
	// 破损赔付
	BrokenCompensatePrice int64 `json:"broken_compensate_price,omitempty" xml:"broken_compensate_price,omitempty"`
	// 首重价格
	BasicWeight int64 `json:"basic_weight,omitempty" xml:"basic_weight,omitempty"`
	// 达成时效
	DeliveryTime int64 `json:"delivery_time,omitempty" xml:"delivery_time,omitempty"`
	// 续重价格
	StepWeight int64 `json:"step_weight,omitempty" xml:"step_weight,omitempty"`
	// 首重价格
	BasicWeightPrice int64 `json:"basic_weight_price,omitempty" xml:"basic_weight_price,omitempty"`
	// 达成率
	AchievingRate int64 `json:"achieving_rate,omitempty" xml:"achieving_rate,omitempty"`
	// 续重价格
	StepWeightPrice int64 `json:"step_weight_price,omitempty" xml:"step_weight_price,omitempty"`
	// 配送资源id
	ResId int64 `json:"res_id,omitempty" xml:"res_id,omitempty"`
	// 丢失赔付价格
	MissingCompensatePrice int64 `json:"missing_compensate_price,omitempty" xml:"missing_compensate_price,omitempty"`
}

var poolThreePlConsignResourceDto = sync.Pool{
	New: func() any {
		return new(ThreePlConsignResourceDto)
	},
}

// GetThreePlConsignResourceDto() 从对象池中获取ThreePlConsignResourceDto
func GetThreePlConsignResourceDto() *ThreePlConsignResourceDto {
	return poolThreePlConsignResourceDto.Get().(*ThreePlConsignResourceDto)
}

// ReleaseThreePlConsignResourceDto 释放ThreePlConsignResourceDto
func ReleaseThreePlConsignResourceDto(v *ThreePlConsignResourceDto) {
	v.ResCode = ""
	v.ResName = ""
	v.BrokenCompensatePrice = 0
	v.BasicWeight = 0
	v.DeliveryTime = 0
	v.StepWeight = 0
	v.BasicWeightPrice = 0
	v.AchievingRate = 0
	v.StepWeightPrice = 0
	v.ResId = 0
	v.MissingCompensatePrice = 0
	poolThreePlConsignResourceDto.Put(v)
}
