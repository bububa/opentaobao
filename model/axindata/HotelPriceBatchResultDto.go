package axindata

import (
	"sync"
)

// HotelPriceBatchResultDto 结构体
type HotelPriceBatchResultDto struct {
	// 酒店的报价结果列表
	HotelPriceDtoList []HotelPriceBatchDto `json:"hotel_price_dto_list,omitempty" xml:"hotel_price_dto_list>hotel_price_batch_dto,omitempty"`
}

var poolHotelPriceBatchResultDto = sync.Pool{
	New: func() any {
		return new(HotelPriceBatchResultDto)
	},
}

// GetHotelPriceBatchResultDto() 从对象池中获取HotelPriceBatchResultDto
func GetHotelPriceBatchResultDto() *HotelPriceBatchResultDto {
	return poolHotelPriceBatchResultDto.Get().(*HotelPriceBatchResultDto)
}

// ReleaseHotelPriceBatchResultDto 释放HotelPriceBatchResultDto
func ReleaseHotelPriceBatchResultDto(v *HotelPriceBatchResultDto) {
	v.HotelPriceDtoList = v.HotelPriceDtoList[:0]
	poolHotelPriceBatchResultDto.Put(v)
}
