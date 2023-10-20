package seaking

import (
	"sync"
)

// ImageTranslateDetailDto 结构体
type ImageTranslateDetailDto struct {
	// 目标语种
	TargetLang string `json:"target_lang,omitempty" xml:"target_lang,omitempty"`
	// 源语种
	SourceLang string `json:"source_lang,omitempty" xml:"source_lang,omitempty"`
	// 图片url
	ImageUrl string `json:"image_url,omitempty" xml:"image_url,omitempty"`
	// 商品所在平台（ae/lazada)）
	Platform string `json:"platform,omitempty" xml:"platform,omitempty"`
	// 商品id
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// 子任务编号，从1开始，必须连续
	Idx int64 `json:"idx,omitempty" xml:"idx,omitempty"`
}

var poolImageTranslateDetailDto = sync.Pool{
	New: func() any {
		return new(ImageTranslateDetailDto)
	},
}

// GetImageTranslateDetailDto() 从对象池中获取ImageTranslateDetailDto
func GetImageTranslateDetailDto() *ImageTranslateDetailDto {
	return poolImageTranslateDetailDto.Get().(*ImageTranslateDetailDto)
}

// ReleaseImageTranslateDetailDto 释放ImageTranslateDetailDto
func ReleaseImageTranslateDetailDto(v *ImageTranslateDetailDto) {
	v.TargetLang = ""
	v.SourceLang = ""
	v.ImageUrl = ""
	v.Platform = ""
	v.ProductId = 0
	v.Idx = 0
	poolImageTranslateDetailDto.Put(v)
}
