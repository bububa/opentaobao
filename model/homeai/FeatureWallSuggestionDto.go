package homeai

import (
	"sync"
)

// FeatureWallSuggestionDto 结构体
type FeatureWallSuggestionDto struct {
	// 墙位置
	FeatureWalls []FeatureWallDto `json:"feature_walls,omitempty" xml:"feature_walls>feature_wall_dto,omitempty"`
	// designid
	DesignId string `json:"design_id,omitempty" xml:"design_id,omitempty"`
}

var poolFeatureWallSuggestionDto = sync.Pool{
	New: func() any {
		return new(FeatureWallSuggestionDto)
	},
}

// GetFeatureWallSuggestionDto() 从对象池中获取FeatureWallSuggestionDto
func GetFeatureWallSuggestionDto() *FeatureWallSuggestionDto {
	return poolFeatureWallSuggestionDto.Get().(*FeatureWallSuggestionDto)
}

// ReleaseFeatureWallSuggestionDto 释放FeatureWallSuggestionDto
func ReleaseFeatureWallSuggestionDto(v *FeatureWallSuggestionDto) {
	v.FeatureWalls = v.FeatureWalls[:0]
	v.DesignId = ""
	poolFeatureWallSuggestionDto.Put(v)
}
