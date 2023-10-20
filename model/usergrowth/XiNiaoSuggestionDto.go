package usergrowth

import (
	"sync"
)

// XiNiaoSuggestionDto 结构体
type XiNiaoSuggestionDto struct {
	// 图片URL
	ImageUrl string `json:"image_url,omitempty" xml:"image_url,omitempty"`
	// 落地页URL
	LandingUrl string `json:"landing_url,omitempty" xml:"landing_url,omitempty"`
	// 曝光上报链接
	ExposureUrl string `json:"exposure_url,omitempty" xml:"exposure_url,omitempty"`
	// 点击上报链接
	ClickUrl string `json:"click_url,omitempty" xml:"click_url,omitempty"`
}

var poolXiNiaoSuggestionDto = sync.Pool{
	New: func() any {
		return new(XiNiaoSuggestionDto)
	},
}

// GetXiNiaoSuggestionDto() 从对象池中获取XiNiaoSuggestionDto
func GetXiNiaoSuggestionDto() *XiNiaoSuggestionDto {
	return poolXiNiaoSuggestionDto.Get().(*XiNiaoSuggestionDto)
}

// ReleaseXiNiaoSuggestionDto 释放XiNiaoSuggestionDto
func ReleaseXiNiaoSuggestionDto(v *XiNiaoSuggestionDto) {
	v.ImageUrl = ""
	v.LandingUrl = ""
	v.ExposureUrl = ""
	v.ClickUrl = ""
	poolXiNiaoSuggestionDto.Put(v)
}
