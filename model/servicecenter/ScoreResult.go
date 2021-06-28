package servicecenter

// ScoreResult 
type ScoreResult struct {

    // 评价id
    
    Id   int64 `json:"id,omitempty" xml:"id,omitempty"`
    

    // 平均分
    
    AvgScore   string `json:"avg_score,omitempty" xml:"avg_score,omitempty"`
    

    // 评论内容
    
    Suggestion   string `json:"suggestion,omitempty" xml:"suggestion,omitempty"`
    

    // 服务code
    
    ServiceCode   string `json:"service_code,omitempty" xml:"service_code,omitempty"`
    

    // 评价人用户昵称
    
    UserNick   string `json:"user_nick,omitempty" xml:"user_nick,omitempty"`
    

    // 评价时间
    
    GmtCreate   string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
    

    // 是否为有效评分 1-有效评分 2-无效评分
    
    IsValid   int64 `json:"is_valid,omitempty" xml:"is_valid,omitempty"`
    

    // 服务规格code
    
    ItemCode   string `json:"item_code,omitempty" xml:"item_code,omitempty"`
    

    // 服务规格名称
    
    ItemName   string `json:"item_name,omitempty" xml:"item_name,omitempty"`
    

    // 是否实际付费 1-实际付费 2-实际未付费
    
    IsPay   int64 `json:"is_pay,omitempty" xml:"is_pay,omitempty"`
    

    // 服务态度评分
    
    AttitudeScore   string `json:"attitude_score,omitempty" xml:"attitude_score,omitempty"`
    

    // 稳定性评分
    
    StabilityScore   string `json:"stability_score,omitempty" xml:"stability_score,omitempty"`
    

    // 易用性评分
    
    EasyuseScore   string `json:"easyuse_score,omitempty" xml:"easyuse_score,omitempty"`
    

    // 专业性评分
    
    ProfScore   string `json:"prof_score,omitempty" xml:"prof_score,omitempty"`
    

    // 交片速度
    
    RapidScore   string `json:"rapid_score,omitempty" xml:"rapid_score,omitempty"`
    

    // 描述相符
    
    MatchedScore   string `json:"matched_score,omitempty" xml:"matched_score,omitempty"`
    

}
