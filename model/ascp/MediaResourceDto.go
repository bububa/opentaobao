package ascp

import (
	"sync"
)

// MediaResourceDto 结构体
type MediaResourceDto struct {
	// 图片/视频名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 图片/视频相对地址
	Path string `json:"path,omitempty" xml:"path,omitempty"`
}

var poolMediaResourceDto = sync.Pool{
	New: func() any {
		return new(MediaResourceDto)
	},
}

// GetMediaResourceDto() 从对象池中获取MediaResourceDto
func GetMediaResourceDto() *MediaResourceDto {
	return poolMediaResourceDto.Get().(*MediaResourceDto)
}

// ReleaseMediaResourceDto 释放MediaResourceDto
func ReleaseMediaResourceDto(v *MediaResourceDto) {
	v.Name = ""
	v.Path = ""
	poolMediaResourceDto.Put(v)
}
