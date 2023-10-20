package alsc

import (
	"sync"
)

// FcUrlDto 结构体
type FcUrlDto struct {
	// 返回url
	ReplaceUrl string `json:"replace_url,omitempty" xml:"replace_url,omitempty"`
}

var poolFcUrlDto = sync.Pool{
	New: func() any {
		return new(FcUrlDto)
	},
}

// GetFcUrlDto() 从对象池中获取FcUrlDto
func GetFcUrlDto() *FcUrlDto {
	return poolFcUrlDto.Get().(*FcUrlDto)
}

// ReleaseFcUrlDto 释放FcUrlDto
func ReleaseFcUrlDto(v *FcUrlDto) {
	v.ReplaceUrl = ""
	poolFcUrlDto.Put(v)
}
