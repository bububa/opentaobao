package usergrowth

import (
	"sync"
)

// SuggestionDto 结构体
type SuggestionDto struct {
	// 标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 副标题
	SubTitle string `json:"sub_title,omitempty" xml:"sub_title,omitempty"`
	// 图标链接
	IconUrl string `json:"icon_url,omitempty" xml:"icon_url,omitempty"`
	// 唤端链接
	DeeplinkUrl string `json:"deeplink_url,omitempty" xml:"deeplink_url,omitempty"`
	// 曝光上报链接
	ExposureUrl string `json:"exposure_url,omitempty" xml:"exposure_url,omitempty"`
	// 点击上报链接
	ClickUrl string `json:"click_url,omitempty" xml:"click_url,omitempty"`
	// 文案内容
	Text string `json:"text,omitempty" xml:"text,omitempty"`
	// 素材id
	MaterialId string `json:"material_id,omitempty" xml:"material_id,omitempty"`
	// 未装端时备用的h5url
	H5Url string `json:"h5_url,omitempty" xml:"h5_url,omitempty"`
	// 主图url
	ImageUrl string `json:"image_url,omitempty" xml:"image_url,omitempty"`
}

var poolSuggestionDto = sync.Pool{
	New: func() any {
		return new(SuggestionDto)
	},
}

// GetSuggestionDto() 从对象池中获取SuggestionDto
func GetSuggestionDto() *SuggestionDto {
	return poolSuggestionDto.Get().(*SuggestionDto)
}

// ReleaseSuggestionDto 释放SuggestionDto
func ReleaseSuggestionDto(v *SuggestionDto) {
	v.Title = ""
	v.SubTitle = ""
	v.IconUrl = ""
	v.DeeplinkUrl = ""
	v.ExposureUrl = ""
	v.ClickUrl = ""
	v.Text = ""
	v.MaterialId = ""
	v.H5Url = ""
	v.ImageUrl = ""
	poolSuggestionDto.Put(v)
}
