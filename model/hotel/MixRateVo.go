package hotel

// MixRateVo 
type MixRateVo struct {

    // 顶数量
    
    AgreeCount   int64 `json:"agree_count,omitempty" xml:"agree_count,omitempty"`
    

    // 所有顶过的所有用户id
    
    AgreeUserIds   []int64 `json:"agree_user_ids,omitempty" xml:"agree_user_ids>int64,omitempty"`
    

    // 业务类型
    
    BizType   int64 `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
    

    // content
    
    Content   string `json:"content,omitempty" xml:"content,omitempty"`
    

    // 踩数量
    
    DisagreeCount   int64 `json:"disagree_count,omitempty" xml:"disagree_count,omitempty"`
    

    // 所有踩过的所有用户id
    
    DisagreeUserIds   []int64 `json:"disagree_user_ids,omitempty" xml:"disagree_user_ids>int64,omitempty"`
    

    // gmtCreate
    
    GmtCreate   string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
    

    // id
    
    Id   int64 `json:"id,omitempty" xml:"id,omitempty"`
    

    // 旅游商品id，对应于ItemRateDO中的travelItemId
    
    ItemId   int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
    

    // 旅游商品信息
    
    ItemInfo   string `json:"item_info,omitempty" xml:"item_info,omitempty"`
    

    // 商品评论ID
    
    ItemRateId   int64 `json:"item_rate_id,omitempty" xml:"item_rate_id,omitempty"`
    

    // 所有商品评论回复
    
    ItemReplies   []ItemRateReplyVo `json:"item_replies,omitempty" xml:"item_replies,omitempty"`
    

    // 点赞数据
    
    Like   *LikeTargetCount `json:"like,omitempty" xml:"like,omitempty"`
    

    // mediaInfo
    
    MediaInfo   string `json:"media_info,omitempty" xml:"media_info,omitempty"`
    

    // 预定id
    
    OrderId   int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
    

    // 预定信息
    
    OrderInfo   string `json:"order_info,omitempty" xml:"order_info,omitempty"`
    

    // 图片链接
    
    PictureUrls   []string `json:"picture_urls,omitempty" xml:"picture_urls>string,omitempty"`
    

    // POI固定文本
    
    PoiStr   string `json:"poi_str,omitempty" xml:"poi_str,omitempty"`
    

    // 跳转链接
    
    RedirectUrl   string `json:"redirect_url,omitempty" xml:"redirect_url,omitempty"`
    

    // 回复数量
    
    ReplyCount   int64 `json:"reply_count,omitempty" xml:"reply_count,omitempty"`
    

    // scoreDetail
    
    ScoreDetail   string `json:"score_detail,omitempty" xml:"score_detail,omitempty"`
    

    // 取件方式:邮寄;套餐类型:7天 4GB 3G流量+无限 2G流量
    
    Sku   string `json:"sku,omitempty" xml:"sku,omitempty"`
    

    // 点评来源
    
    Source   int64 `json:"source,omitempty" xml:"source,omitempty"`
    

    // 点评来源名称
    
    SourceTypeName   string `json:"source_type_name,omitempty" xml:"source_type_name,omitempty"`
    

    // 以上点评来自TripAdvisor.
    
    SplitLineContent   string `json:"split_line_content,omitempty" xml:"split_line_content,omitempty"`
    

    // 状态
    
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`
    

    // 标签
    
    TagInfo   string `json:"tag_info,omitempty" xml:"tag_info,omitempty"`
    

    // 标题
    
    Title   string `json:"title,omitempty" xml:"title,omitempty"`
    

    // 总分
    
    TotalScore   int64 `json:"total_score,omitempty" xml:"total_score,omitempty"`
    

    // 行程名称
    
    TravelName   string `json:"travel_name,omitempty" xml:"travel_name,omitempty"`
    

    // 航旅标准商品子类型id
    
    TravelSubItemId   int64 `json:"travel_sub_item_id,omitempty" xml:"travel_sub_item_id,omitempty"`
    

    // 航旅标准商品子类型id
    
    TravelSubItemInfo   string `json:"travel_sub_item_info,omitempty" xml:"travel_sub_item_info,omitempty"`
    

    // 行程ID
    
    TripGuidId   int64 `json:"trip_guid_id,omitempty" xml:"trip_guid_id,omitempty"`
    

    // ttid
    
    Ttid   string `json:"ttid,omitempty" xml:"ttid,omitempty"`
    

    // 用户头像
    
    UserIcon   string `json:"user_icon,omitempty" xml:"user_icon,omitempty"`
    

    // 脱敏后的用户id
    
    UserId   int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
    

    // 脱敏后的用户名字
    
    UserNick   string `json:"user_nick,omitempty" xml:"user_nick,omitempty"`
    

    // 用户星级
    
    UserStar   int64 `json:"user_star,omitempty" xml:"user_star,omitempty"`
    

}
