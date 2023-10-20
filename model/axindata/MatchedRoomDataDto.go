package axindata

import (
	"sync"
)

// MatchedRoomDataDto 结构体
type MatchedRoomDataDto struct {
	// 分值,越高可信度越高
	Score string `json:"score,omitempty" xml:"score,omitempty"`
	// 匹配结果标准房型id
	Srid int64 `json:"srid,omitempty" xml:"srid,omitempty"`
}

var poolMatchedRoomDataDto = sync.Pool{
	New: func() any {
		return new(MatchedRoomDataDto)
	},
}

// GetMatchedRoomDataDto() 从对象池中获取MatchedRoomDataDto
func GetMatchedRoomDataDto() *MatchedRoomDataDto {
	return poolMatchedRoomDataDto.Get().(*MatchedRoomDataDto)
}

// ReleaseMatchedRoomDataDto 释放MatchedRoomDataDto
func ReleaseMatchedRoomDataDto(v *MatchedRoomDataDto) {
	v.Score = ""
	v.Srid = 0
	poolMatchedRoomDataDto.Put(v)
}
