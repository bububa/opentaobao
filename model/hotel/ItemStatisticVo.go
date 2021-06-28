package hotel

// ItemStatisticVo 
type ItemStatisticVo struct {

    // 最佳得分项
    
    BestItem   string `json:"best_item,omitempty" xml:"best_item,omitempty"`
    

    // 五分制标记, 1为五分制， 0为十分制
    
    IsFiveGrade   int64 `json:"is_five_grade,omitempty" xml:"is_five_grade,omitempty"`
    

    // 评论总数
    
    RateCnt   int64 `json:"rate_cnt,omitempty" xml:"rate_cnt,omitempty"`
    

    // 有图的评论总数
    
    RatePicCnt   int64 `json:"rate_pic_cnt,omitempty" xml:"rate_pic_cnt,omitempty"`
    

    // 推荐率
    
    RecommendStr   string `json:"recommend_str,omitempty" xml:"recommend_str,omitempty"`
    

    // tab信息
    
    RoomTabInfos   []TabInfo `json:"room_tab_infos,omitempty" xml:"room_tab_infos,omitempty"`
    

    // 评分描述： 非常好
    
    ScoreDesc   string `json:"score_desc,omitempty" xml:"score_desc,omitempty"`
    

    // 评分详情，json格式
    
    ScoreDetail   string `json:"score_detail,omitempty" xml:"score_detail,omitempty"`
    

    // 不同分数的数量
    
    ScoreInfos   []ScoreInfo `json:"score_infos,omitempty" xml:"score_infos,omitempty"`
    

    // 评分星级
    
    ScoreLevel   int64 `json:"score_level,omitempty" xml:"score_level,omitempty"`
    

    // source来源 0自采 1共享 21agoda 22艺龙 23tripAdvisor
    
    Source   int64 `json:"source,omitempty" xml:"source,omitempty"`
    

    // tab信息
    
    TabInfos   []TabInfo `json:"tab_infos,omitempty" xml:"tab_infos,omitempty"`
    

    // 热词显示的行数
    
    TabShowLines   int64 `json:"tab_show_lines,omitempty" xml:"tab_show_lines,omitempty"`
    

    // 总评分
    
    TotalScore   string `json:"total_score,omitempty" xml:"total_score,omitempty"`
    

    // 旅游商品信息id
    
    TravelItemId   int64 `json:"travel_item_id,omitempty" xml:"travel_item_id,omitempty"`
    

    // 旅游商品信息
    
    TravelItemInfo   string `json:"travel_item_info,omitempty" xml:"travel_item_info,omitempty"`
    

    // tripAdv评论数
    
    TripAdvateCnt   int64 `json:"trip_advate_cnt,omitempty" xml:"trip_advate_cnt,omitempty"`
    

}
