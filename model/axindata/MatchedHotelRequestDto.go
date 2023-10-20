package axindata

import (
	"sync"
)

// MatchedHotelRequestDto 结构体
type MatchedHotelRequestDto struct {
	// 要匹配的酒店参数列表
	HotelMatchDTOList []HotelMatchDto `json:"hotel_match_d_t_o_list,omitempty" xml:"hotel_match_d_t_o_list>hotel_match_dto,omitempty"`
	// 分销商id
	DistributorTid int64 `json:"distributor_tid,omitempty" xml:"distributor_tid,omitempty"`
}

var poolMatchedHotelRequestDto = sync.Pool{
	New: func() any {
		return new(MatchedHotelRequestDto)
	},
}

// GetMatchedHotelRequestDto() 从对象池中获取MatchedHotelRequestDto
func GetMatchedHotelRequestDto() *MatchedHotelRequestDto {
	return poolMatchedHotelRequestDto.Get().(*MatchedHotelRequestDto)
}

// ReleaseMatchedHotelRequestDto 释放MatchedHotelRequestDto
func ReleaseMatchedHotelRequestDto(v *MatchedHotelRequestDto) {
	v.HotelMatchDTOList = v.HotelMatchDTOList[:0]
	v.DistributorTid = 0
	poolMatchedHotelRequestDto.Put(v)
}
