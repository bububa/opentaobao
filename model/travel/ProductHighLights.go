package travel

// ProductHighLights 结构体
type ProductHighLights struct {
	// 产品亮点标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 产品亮点描述
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 产品亮点图片
	PicUrls []string `json:"pic_urls,omitempty" xml:"pic_urls>string,omitempty"`
}
