package traderate

// TmallRateInfo 
type TmallRateInfo struct {

    // 表示评价者的昵称
    
    UserNick   string `json:"user_nick,omitempty" xml:"user_nick,omitempty"`
    

    // 评价内容
    
    Content   string `json:"content,omitempty" xml:"content,omitempty"`
    

    // 评价时间
    
    CommentTime   string `json:"comment_time,omitempty" xml:"comment_time,omitempty"`
    

    // 原始评价对应的标签列表
    
    Tags   []TmallRateTag `json:"tags,omitempty" xml:"tags,omitempty"`
    

    // 原始评价是否含有负向标签
    
    HasNegtv   bool `json:"has_negtv,omitempty" xml:"has_negtv,omitempty"`
    

    // 追加评价内容
    
    AppendContent   string `json:"append_content,omitempty" xml:"append_content,omitempty"`
    

    // 追加评价时间
    
    AppendTime   string `json:"append_time,omitempty" xml:"append_time,omitempty"`
    

    // 追加评价中带有的语义标签列表
    
    AppendTags   []TmallRateTag `json:"append_tags,omitempty" xml:"append_tags,omitempty"`
    

    // 追评中是否含有负向标签
    
    AppendHasNegtv   bool `json:"append_has_negtv,omitempty" xml:"append_has_negtv,omitempty"`
    

}
