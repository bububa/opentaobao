package usergrowth

// ExtraDto 结构体
type ExtraDto struct {
	// 拓展曝光上报链接, 可以用于榜单曝光
	ExposureUrl []string `json:"exposure_url,omitempty" xml:"exposure_url>string,omitempty"`
	// 拓展点击上报链接, 可以用于“点击更多”上报
	ClickUrl []string `json:"click_url,omitempty" xml:"click_url>string,omitempty"`
	// App跳转链接
	DeeplinkUrl string `json:"deeplink_url,omitempty" xml:"deeplink_url,omitempty"`
	// 未安装淘宝时 h5 页面跳
	H5Url string `json:"h5_url,omitempty" xml:"h5_url,omitempty"`
}
