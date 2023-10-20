package util

import (
	"sync"
)

// ShortUrlDto 结构体
type ShortUrlDto struct {
	// shortUrl
	ShortUrl string `json:"short_url,omitempty" xml:"short_url,omitempty"`
}

var poolShortUrlDto = sync.Pool{
	New: func() any {
		return new(ShortUrlDto)
	},
}

// GetShortUrlDto() 从对象池中获取ShortUrlDto
func GetShortUrlDto() *ShortUrlDto {
	return poolShortUrlDto.Get().(*ShortUrlDto)
}

// ReleaseShortUrlDto 释放ShortUrlDto
func ReleaseShortUrlDto(v *ShortUrlDto) {
	v.ShortUrl = ""
	poolShortUrlDto.Put(v)
}
