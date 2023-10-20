package alihouse

import (
	"sync"
)

// FeatureDto 结构体
type FeatureDto struct {
	// 1
	Key string `json:"key,omitempty" xml:"key,omitempty"`
	// 1
	Value string `json:"value,omitempty" xml:"value,omitempty"`
	// 1
	Source string `json:"source,omitempty" xml:"source,omitempty"`
}

var poolFeatureDto = sync.Pool{
	New: func() any {
		return new(FeatureDto)
	},
}

// GetFeatureDto() 从对象池中获取FeatureDto
func GetFeatureDto() *FeatureDto {
	return poolFeatureDto.Get().(*FeatureDto)
}

// ReleaseFeatureDto 释放FeatureDto
func ReleaseFeatureDto(v *FeatureDto) {
	v.Key = ""
	v.Value = ""
	v.Source = ""
	poolFeatureDto.Put(v)
}
