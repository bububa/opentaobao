package aesolution

import (
	"sync"
)

// MarketImageDto 结构体
type MarketImageDto struct {
	// The image url needs to be obtained via uploading the image through Aliexpress API: aliexpress.photobank.redefining.uploadimageforsdk(https://developers.aliexpress.com/en/doc.htm?docId=30186&amp;docType=2)
	ImageUrl string `json:"image_url,omitempty" xml:"image_url,omitempty"`
	// 1 represents 3:4 rectangle(resolution at least 750*1000) image while 2 represents 1:1 square image(Resolution at least 800*800).
	ImageType string `json:"image_type,omitempty" xml:"image_type,omitempty"`
}

var poolMarketImageDto = sync.Pool{
	New: func() any {
		return new(MarketImageDto)
	},
}

// GetMarketImageDto() 从对象池中获取MarketImageDto
func GetMarketImageDto() *MarketImageDto {
	return poolMarketImageDto.Get().(*MarketImageDto)
}

// ReleaseMarketImageDto 释放MarketImageDto
func ReleaseMarketImageDto(v *MarketImageDto) {
	v.ImageUrl = ""
	v.ImageType = ""
	poolMarketImageDto.Put(v)
}
