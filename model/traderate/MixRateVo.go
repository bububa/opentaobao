package traderate

// MixRateVo 
/* model for simplify = false
type MixRateVo struct {

    // 正文内容
    
    Content   string `json:"content,omitempty"`
    

    // 创建时间
    
    GmtCreate   string `json:"gmt_create,omitempty"`
    

    // 图片信息
    
    PictureUrls  struct {
        String  []string `json:"string,omitempty"`
    } `json:"picture_urls,omitempty"`
    

    // 正文摘要
    
    Title   string `json:"title,omitempty"`
    

    // 用户头像
    
    UserIcon   string `json:"user_icon,omitempty"`
    

    // 用户昵称
    
    UserNick   string `json:"user_nick,omitempty"`
    

    // 总评分
    
    TotalScore   int64 `json:"total_score,omitempty"`
    

}
*/

// MixRateVo 
type MixRateVo struct {

    // 正文内容
    Content   string `json:"content,omitempty"`

    // 创建时间
    GmtCreate   string `json:"gmt_create,omitempty"`

    // 图片信息
    PictureUrls   []string `json:"picture_urls,omitempty"`

    // 正文摘要
    Title   string `json:"title,omitempty"`

    // 用户头像
    UserIcon   string `json:"user_icon,omitempty"`

    // 用户昵称
    UserNick   string `json:"user_nick,omitempty"`

    // 总评分
    TotalScore   int64 `json:"total_score,omitempty"`

}
