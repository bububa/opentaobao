package axindata

import (
	"sync"
)

// MatchedHotelBatchResultDto 结构体
type MatchedHotelBatchResultDto struct {
	// 酒店匹配结果列表
	MatchedHotelResultList []MatchedHotelResultDto `json:"matched_hotel_result_list,omitempty" xml:"matched_hotel_result_list>matched_hotel_result_dto,omitempty"`
}

var poolMatchedHotelBatchResultDto = sync.Pool{
	New: func() any {
		return new(MatchedHotelBatchResultDto)
	},
}

// GetMatchedHotelBatchResultDto() 从对象池中获取MatchedHotelBatchResultDto
func GetMatchedHotelBatchResultDto() *MatchedHotelBatchResultDto {
	return poolMatchedHotelBatchResultDto.Get().(*MatchedHotelBatchResultDto)
}

// ReleaseMatchedHotelBatchResultDto 释放MatchedHotelBatchResultDto
func ReleaseMatchedHotelBatchResultDto(v *MatchedHotelBatchResultDto) {
	v.MatchedHotelResultList = v.MatchedHotelResultList[:0]
	poolMatchedHotelBatchResultDto.Put(v)
}
