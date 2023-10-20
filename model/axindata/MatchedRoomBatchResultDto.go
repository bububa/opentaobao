package axindata

import (
	"sync"
)

// MatchedRoomBatchResultDto 结构体
type MatchedRoomBatchResultDto struct {
	// 每个房型匹配结果封装对象
	MatchedRoomResultList []MatchedRoomResultDto `json:"matched_room_result_list,omitempty" xml:"matched_room_result_list>matched_room_result_dto,omitempty"`
}

var poolMatchedRoomBatchResultDto = sync.Pool{
	New: func() any {
		return new(MatchedRoomBatchResultDto)
	},
}

// GetMatchedRoomBatchResultDto() 从对象池中获取MatchedRoomBatchResultDto
func GetMatchedRoomBatchResultDto() *MatchedRoomBatchResultDto {
	return poolMatchedRoomBatchResultDto.Get().(*MatchedRoomBatchResultDto)
}

// ReleaseMatchedRoomBatchResultDto 释放MatchedRoomBatchResultDto
func ReleaseMatchedRoomBatchResultDto(v *MatchedRoomBatchResultDto) {
	v.MatchedRoomResultList = v.MatchedRoomResultList[:0]
	poolMatchedRoomBatchResultDto.Put(v)
}
