package axindata

import (
	"sync"
)

// MatchedHotelDataDto 结构体
type MatchedHotelDataDto struct {
	// 分值
	Score string `json:"score,omitempty" xml:"score,omitempty"`
	// 匹配的标准酒店id
	Shid int64 `json:"shid,omitempty" xml:"shid,omitempty"`
}

var poolMatchedHotelDataDto = sync.Pool{
	New: func() any {
		return new(MatchedHotelDataDto)
	},
}

// GetMatchedHotelDataDto() 从对象池中获取MatchedHotelDataDto
func GetMatchedHotelDataDto() *MatchedHotelDataDto {
	return poolMatchedHotelDataDto.Get().(*MatchedHotelDataDto)
}

// ReleaseMatchedHotelDataDto 释放MatchedHotelDataDto
func ReleaseMatchedHotelDataDto(v *MatchedHotelDataDto) {
	v.Score = ""
	v.Shid = 0
	poolMatchedHotelDataDto.Put(v)
}
