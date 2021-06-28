package simba

// KeywordQscore 
/* model for simplify = false
type KeywordQscore struct {

    // 主人昵称
    
    Nick   string `json:"nick,omitempty"`
    

    // 推广计划id
    
    CampaignId   int64 `json:"campaign_id,omitempty"`
    

    // 推广组id
    
    AdgroupId   int64 `json:"adgroup_id,omitempty"`
    

    // 关键词id
    
    KeywordId   int64 `json:"keyword_id,omitempty"`
    

    // 关键词
    
    Word   string `json:"word,omitempty"`
    

    // 质量总分
    
    Qscore   string `json:"qscore,omitempty"`
    

    // 相关性
    
    ReleScore   string `json:"rele_score,omitempty"`
    

    // 买家体验分
    
    CvrScore   string `json:"cvr_score,omitempty"`
    

    // 基础分
    
    CustScore   string `json:"cust_score,omitempty"`
    

    // 创意得分
    
    CreativeScore   string `json:"creative_score,omitempty"`
    

}
*/

// KeywordQscore 
type KeywordQscore struct {

    // 主人昵称
    Nick   string `json:"nick,omitempty"`

    // 推广计划id
    CampaignId   int64 `json:"campaign_id,omitempty"`

    // 推广组id
    AdgroupId   int64 `json:"adgroup_id,omitempty"`

    // 关键词id
    KeywordId   int64 `json:"keyword_id,omitempty"`

    // 关键词
    Word   string `json:"word,omitempty"`

    // 质量总分
    Qscore   string `json:"qscore,omitempty"`

    // 相关性
    ReleScore   string `json:"rele_score,omitempty"`

    // 买家体验分
    CvrScore   string `json:"cvr_score,omitempty"`

    // 基础分
    CustScore   string `json:"cust_score,omitempty"`

    // 创意得分
    CreativeScore   string `json:"creative_score,omitempty"`

}
