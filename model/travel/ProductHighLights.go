package travel

// ProductHighLights 
type ProductHighLights struct {

    // 产品亮点标题
    Title   string `json:"title,omitempty"`

    // 产品亮点描述
    Desc   string `json:"desc,omitempty"`

    // 产品亮点图片
    PicUrls   []String `json:"pic_urls,omitempty"`

}
