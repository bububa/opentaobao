package axindata

import (
	"sync"
)

// MatchedRoomRequestDto 结构体
type MatchedRoomRequestDto struct {
	// 要匹配的房型参数列表
	HotelRoomMatchDTOList []HotelRoomMatchDto `json:"hotel_room_match_d_t_o_list,omitempty" xml:"hotel_room_match_d_t_o_list>hotel_room_match_dto,omitempty"`
	// 分销商淘宝账号id
	DistributorTid int64 `json:"distributor_tid,omitempty" xml:"distributor_tid,omitempty"`
}

var poolMatchedRoomRequestDto = sync.Pool{
	New: func() any {
		return new(MatchedRoomRequestDto)
	},
}

// GetMatchedRoomRequestDto() 从对象池中获取MatchedRoomRequestDto
func GetMatchedRoomRequestDto() *MatchedRoomRequestDto {
	return poolMatchedRoomRequestDto.Get().(*MatchedRoomRequestDto)
}

// ReleaseMatchedRoomRequestDto 释放MatchedRoomRequestDto
func ReleaseMatchedRoomRequestDto(v *MatchedRoomRequestDto) {
	v.HotelRoomMatchDTOList = v.HotelRoomMatchDTOList[:0]
	v.DistributorTid = 0
	poolMatchedRoomRequestDto.Put(v)
}
