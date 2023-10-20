package axindata

import (
	"sync"
)

// FscProductCancelPolicyApiDto 结构体
type FscProductCancelPolicyApiDto struct {
	// 扣除明细
	DeductList []FscProductCancelDeductApiDto `json:"deduct_list,omitempty" xml:"deduct_list>fsc_product_cancel_deduct_api_dto,omitempty"`
	// 取消类型 NORMAL: 平日取消政策 VACATION: 节假日取消政策（法定节假日，不含周末）
	CancelType string `json:"cancel_type,omitempty" xml:"cancel_type,omitempty"`
	// 提前小时分钟 格式：HH:MM，10:30代表10小时30分钟；总提前时间为提前天数加上提前小时分钟，如：1天10小时30分钟，不填默认为 00:00
	AheadHourMinute string `json:"ahead_hour_minute,omitempty" xml:"ahead_hour_minute,omitempty"`
	// 提前天数 1: 代表1天
	AheadDays int64 `json:"ahead_days,omitempty" xml:"ahead_days,omitempty"`
}

var poolFscProductCancelPolicyApiDto = sync.Pool{
	New: func() any {
		return new(FscProductCancelPolicyApiDto)
	},
}

// GetFscProductCancelPolicyApiDto() 从对象池中获取FscProductCancelPolicyApiDto
func GetFscProductCancelPolicyApiDto() *FscProductCancelPolicyApiDto {
	return poolFscProductCancelPolicyApiDto.Get().(*FscProductCancelPolicyApiDto)
}

// ReleaseFscProductCancelPolicyApiDto 释放FscProductCancelPolicyApiDto
func ReleaseFscProductCancelPolicyApiDto(v *FscProductCancelPolicyApiDto) {
	v.DeductList = v.DeductList[:0]
	v.CancelType = ""
	v.AheadHourMinute = ""
	v.AheadDays = 0
	poolFscProductCancelPolicyApiDto.Put(v)
}
