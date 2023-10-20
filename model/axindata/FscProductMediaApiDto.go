package axindata

import (
	"sync"
)

// FscProductMediaApiDto 结构体
type FscProductMediaApiDto struct {
	// 供应商多媒体编码
	MediaCode string `json:"media_code,omitempty" xml:"media_code,omitempty"`
	// 多媒体下载链接
	MediaUrl string `json:"media_url,omitempty" xml:"media_url,omitempty"`
	// 多媒体类型 1:图片 2:视频
	MediaType int64 `json:"media_type,omitempty" xml:"media_type,omitempty"`
	// 多媒体排序值
	Sort int64 `json:"sort,omitempty" xml:"sort,omitempty"`
}

var poolFscProductMediaApiDto = sync.Pool{
	New: func() any {
		return new(FscProductMediaApiDto)
	},
}

// GetFscProductMediaApiDto() 从对象池中获取FscProductMediaApiDto
func GetFscProductMediaApiDto() *FscProductMediaApiDto {
	return poolFscProductMediaApiDto.Get().(*FscProductMediaApiDto)
}

// ReleaseFscProductMediaApiDto 释放FscProductMediaApiDto
func ReleaseFscProductMediaApiDto(v *FscProductMediaApiDto) {
	v.MediaCode = ""
	v.MediaUrl = ""
	v.MediaType = 0
	v.Sort = 0
	poolFscProductMediaApiDto.Put(v)
}
