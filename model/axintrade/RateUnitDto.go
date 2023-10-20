package axintrade

import (
	"sync"
)

// RateUnitDto 结构体
type RateUnitDto struct {
	// 房间价格
	DailyPriceInfoList []DailyPriceInfoDto `json:"daily_price_info_list,omitempty" xml:"daily_price_info_list>daily_price_info_dto,omitempty"`
}

var poolRateUnitDto = sync.Pool{
	New: func() any {
		return new(RateUnitDto)
	},
}

// GetRateUnitDto() 从对象池中获取RateUnitDto
func GetRateUnitDto() *RateUnitDto {
	return poolRateUnitDto.Get().(*RateUnitDto)
}

// ReleaseRateUnitDto 释放RateUnitDto
func ReleaseRateUnitDto(v *RateUnitDto) {
	v.DailyPriceInfoList = v.DailyPriceInfoList[:0]
	poolRateUnitDto.Put(v)
}
