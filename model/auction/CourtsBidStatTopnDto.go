package auction

import (
	"sync"
)

// CourtsBidStatTopnDto 结构体
type CourtsBidStatTopnDto struct {
	// 法院名称
	CourtName string `json:"court_name,omitempty" xml:"court_name,omitempty"`
	// 排名
	Rank int64 `json:"rank,omitempty" xml:"rank,omitempty"`
	// 成交价（成交标的）
	HammerPrice int64 `json:"hammer_price,omitempty" xml:"hammer_price,omitempty"`
	// 发拍件数（去重）
	PublishCountDist int64 `json:"publish_count_dist,omitempty" xml:"publish_count_dist,omitempty"`
}

var poolCourtsBidStatTopnDto = sync.Pool{
	New: func() any {
		return new(CourtsBidStatTopnDto)
	},
}

// GetCourtsBidStatTopnDto() 从对象池中获取CourtsBidStatTopnDto
func GetCourtsBidStatTopnDto() *CourtsBidStatTopnDto {
	return poolCourtsBidStatTopnDto.Get().(*CourtsBidStatTopnDto)
}

// ReleaseCourtsBidStatTopnDto 释放CourtsBidStatTopnDto
func ReleaseCourtsBidStatTopnDto(v *CourtsBidStatTopnDto) {
	v.CourtName = ""
	v.Rank = 0
	v.HammerPrice = 0
	v.PublishCountDist = 0
	poolCourtsBidStatTopnDto.Put(v)
}
