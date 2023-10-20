package alihouse

import (
	"sync"
)

// PointsFactorDto 结构体
type PointsFactorDto struct {
	// 经纪人积分因素
	PointFactor string `json:"point_factor,omitempty" xml:"point_factor,omitempty"`
	// 经纪人积分分值
	PointScore int64 `json:"point_score,omitempty" xml:"point_score,omitempty"`
}

var poolPointsFactorDto = sync.Pool{
	New: func() any {
		return new(PointsFactorDto)
	},
}

// GetPointsFactorDto() 从对象池中获取PointsFactorDto
func GetPointsFactorDto() *PointsFactorDto {
	return poolPointsFactorDto.Get().(*PointsFactorDto)
}

// ReleasePointsFactorDto 释放PointsFactorDto
func ReleasePointsFactorDto(v *PointsFactorDto) {
	v.PointFactor = ""
	v.PointScore = 0
	poolPointsFactorDto.Put(v)
}
