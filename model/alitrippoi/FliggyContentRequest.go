package alitrippoi

// FliggyContentRequest 
type FliggyContentRequest struct {

    // 内容摘要
    
    Summary   string `json:"summary,omitempty" xml:"summary,omitempty"`
    

    // 发布时间
    
    PublishDate   string `json:"publish_date,omitempty" xml:"publish_date,omitempty"`
    

    // 发布者用户名
    
    UserName   string `json:"user_name,omitempty" xml:"user_name,omitempty"`
    

    // 标题
    
    Title   string `json:"title,omitempty" xml:"title,omitempty"`
    

    // 城市信息
    
    Citys   []string `json:"citys,omitempty" xml:"citys>string,omitempty"`
    

    // 发布者数字id
    
    UserId   string `json:"user_id,omitempty" xml:"user_id,omitempty"`
    

    // 父文章ID
    
    ParentId   string `json:"parent_id,omitempty" xml:"parent_id,omitempty"`
    

    // 内容正文
    
    Content   string `json:"content,omitempty" xml:"content,omitempty"`
    

    // 标签列表
    
    TagList   []string `json:"tag_list,omitempty" xml:"tag_list>string,omitempty"`
    

    // 其它特征
    
    Feature   string `json:"feature,omitempty" xml:"feature,omitempty"`
    

    // 文章类型
    
    ArticleType   string `json:"article_type,omitempty" xml:"article_type,omitempty"`
    

    // 文章ID
    
    Id   string `json:"id,omitempty" xml:"id,omitempty"`
    

    // 分类
    
    Category   string `json:"category,omitempty" xml:"category,omitempty"`
    

    // 图片列表
    
    ImgList   []string `json:"img_list,omitempty" xml:"img_list>string,omitempty"`
    

    // 视频封面url
    
    VideoCoverUrl   string `json:"video_cover_url,omitempty" xml:"video_cover_url,omitempty"`
    

    // 视频url
    
    VideoUrl   string `json:"video_url,omitempty" xml:"video_url,omitempty"`
    

    // 用户icon
    
    UserIcon   string `json:"user_icon,omitempty" xml:"user_icon,omitempty"`
    

}
