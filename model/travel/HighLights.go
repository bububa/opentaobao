package travel

// HighLights 
type HighLights struct {

    // 亮点标题
    Title   string `json:"title,omitempty"`

    // 亮点描述
    Desc   string `json:"desc,omitempty"`

    // 图片列表
    PicUrls   []String `json:"pic_urls,omitempty"`

}
