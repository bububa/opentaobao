package usergrowth

// BuzzwordDto 结构体
type BuzzwordDto struct {
	// 点击上报链接
	ClickUrl []string `json:"click_url,omitempty" xml:"click_url>string,omitempty"`
	// 词条曝光链接
	ExposureUrl []string `json:"exposure_url,omitempty" xml:"exposure_url>string,omitempty"`
	// 词条
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 标签，可选，热、新等
	Tag string `json:"tag,omitempty" xml:"tag,omitempty"`
	// 标签对应的图片 url
	TagUrl string `json:"tag_url,omitempty" xml:"tag_url,omitempty"`
	// App跳转链接
	DeeplinkUrl string `json:"deeplink_url,omitempty" xml:"deeplink_url,omitempty"`
	// 未安装淘宝时 h5 页面跳
	H5Url string `json:"h5_url,omitempty" xml:"h5_url,omitempty"`
	// 展示图片
	DisplayImageUrl string `json:"display_image_url,omitempty" xml:"display_image_url,omitempty"`
	// 热度搜索文字
	HeatSearchText string `json:"heat_search_text,omitempty" xml:"heat_search_text,omitempty"`
	// 副标题
	SubTitle string `json:"sub_title,omitempty" xml:"sub_title,omitempty"`
	// 热度
	Popularity int64 `json:"popularity,omitempty" xml:"popularity,omitempty"`
	// 词条失效时间(秒级时间戳)
	InvalidTime int64 `json:"invalid_time,omitempty" xml:"invalid_time,omitempty"`
}
