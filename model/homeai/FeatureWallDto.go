package homeai

import (
	"sync"
)

// FeatureWallDto 结构体
type FeatureWallDto struct {
	// 方向四元数
	Fronts []string `json:"fronts,omitempty" xml:"fronts>string,omitempty"`
	// 坐标
	Pos []string `json:"pos,omitempty" xml:"pos>string,omitempty"`
	// roomid
	Room string `json:"room,omitempty" xml:"room,omitempty"`
	// 硬装类型
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 包围盒
	BoundingBox *BoundingBoxDto `json:"bounding_box,omitempty" xml:"bounding_box,omitempty"`
}

var poolFeatureWallDto = sync.Pool{
	New: func() any {
		return new(FeatureWallDto)
	},
}

// GetFeatureWallDto() 从对象池中获取FeatureWallDto
func GetFeatureWallDto() *FeatureWallDto {
	return poolFeatureWallDto.Get().(*FeatureWallDto)
}

// ReleaseFeatureWallDto 释放FeatureWallDto
func ReleaseFeatureWallDto(v *FeatureWallDto) {
	v.Fronts = v.Fronts[:0]
	v.Pos = v.Pos[:0]
	v.Room = ""
	v.Type = ""
	v.BoundingBox = nil
	poolFeatureWallDto.Put(v)
}
