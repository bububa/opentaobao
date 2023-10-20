package tbk

import (
	"sync"
)

// StepRateDto 结构体
type StepRateDto struct {
	// 前N件佣金结束时间， 当前N件佣金 失效，本字段置空
	TopnEndTime int64 `json:"topn_end_time,omitempty" xml:"topn_end_time,omitempty"`
	// 前N件佣金开始时间，当前N件佣金失效，本字段置空
	TopnStartTime int64 `json:"topn_start_time,omitempty" xml:"topn_start_time,omitempty"`
	// 前N件剩余库存，当前N件佣金失效，本字段置空
	TopnQuantity int64 `json:"topn_quantity,omitempty" xml:"topn_quantity,omitempty"`
	// 前N件初始总库存，当前N件佣金失效，本字段置空（失效：任务完成、时间结束、商品下架）
	TopnTotalCount int64 `json:"topn_total_count,omitempty" xml:"topn_total_count,omitempty"`
}

var poolStepRateDto = sync.Pool{
	New: func() any {
		return new(StepRateDto)
	},
}

// GetStepRateDto() 从对象池中获取StepRateDto
func GetStepRateDto() *StepRateDto {
	return poolStepRateDto.Get().(*StepRateDto)
}

// ReleaseStepRateDto 释放StepRateDto
func ReleaseStepRateDto(v *StepRateDto) {
	v.TopnEndTime = 0
	v.TopnStartTime = 0
	v.TopnQuantity = 0
	v.TopnTotalCount = 0
	poolStepRateDto.Put(v)
}
