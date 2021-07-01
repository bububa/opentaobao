package travel

// HighLights 结构体
type HighLights struct {
	// 亮点标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 亮点描述
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 图片列表
	PicUrls []string `json:"pic_urls,omitempty" xml:"pic_urls>string,omitempty"`
}
