package usergrowth

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
