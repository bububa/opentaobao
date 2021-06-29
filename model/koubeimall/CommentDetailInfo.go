package koubeimall

// CommentDetailInfo 
type CommentDetailInfo struct {

    // 图片列表
    
    CommentImageList   []string `json:"comment_image_list,omitempty" xml:"comment_image_list>string,omitempty"`
    

    // 评分
    
    CommentScore   int64 `json:"comment_score,omitempty" xml:"comment_score,omitempty"`
    

    // 评论内容
    
    CommentContent   string `json:"comment_content,omitempty" xml:"comment_content,omitempty"`
    

    // 评价时间
    
    CommentTime   string `json:"comment_time,omitempty" xml:"comment_time,omitempty"`
    

}
