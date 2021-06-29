package alihealthcrm

// Contents 
type Contents struct {

    // 文章标题
    
    Title   string `json:"title,omitempty" xml:"title,omitempty"`
    

    // 文章标签
    
    Tags   string `json:"tags,omitempty" xml:"tags,omitempty"`
    

    // 图片
    
    PhotoUrl   string `json:"photo_url,omitempty" xml:"photo_url,omitempty"`
    

    // 内容
    
    Content   string `json:"content,omitempty" xml:"content,omitempty"`
    

    // 链接
    
    LinkUrl   string `json:"link_url,omitempty" xml:"link_url,omitempty"`
    

    // 商品计数
    
    ItemCount   int64 `json:"item_count,omitempty" xml:"item_count,omitempty"`
    

    // 商品列表
    
    Items   []string `json:"items,omitempty" xml:"items>string,omitempty"`
    

    // publishTime
    
    PublishTime   string `json:"publish_time,omitempty" xml:"publish_time,omitempty"`
    

}
