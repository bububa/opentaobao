package alihealthcrm

// Content 
type Content struct {
    // 标题
    Title   string `json:"title,omitempty" xml:"title,omitempty"`
    // 跳转链接
    Link   string `json:"link,omitempty" xml:"link,omitempty"`
    // 卡片code 0=首页 ， 2544 =减重卡片，41596=孕产，64999=育儿，5122=备孕
    TagCode   string `json:"tag_code,omitempty" xml:"tag_code,omitempty"`
    // 卡片状态
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`
    // 卡片已开启数
    OpenCount   int64 `json:"open_count,omitempty" xml:"open_count,omitempty"`
    // 点击后的图标
    ImgUrlClicked   string `json:"img_url_clicked,omitempty" xml:"img_url_clicked,omitempty"`
    // 未点击时的图标
    ImgUrl   string `json:"img_url,omitempty" xml:"img_url,omitempty"`
    // 字体颜色，第一个为未点击时的，第二个为点击后的
    Color   string `json:"color,omitempty" xml:"color,omitempty"`
}
