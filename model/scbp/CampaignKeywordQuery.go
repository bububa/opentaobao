package scbp

// CampaignKeywordQuery 
type CampaignKeywordQuery struct {

    // 普通词
    
    NormWord   string `json:"norm_word,omitempty" xml:"norm_word,omitempty"`
    

    // 词id集合
    
    WordIdList   []int64 `json:"word_id_list,omitempty" xml:"word_id_list>int64,omitempty"`
    

    // 计划id
    
    CampaignId   int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
    

    // 状态
    
    OnlineStatus   int64 `json:"online_status,omitempty" xml:"online_status,omitempty"`
    

    // 类型:base(基础信息)，star(基础信息加星级),full（所有信息），不传默认查所有信息
    
    Type   string `json:"type,omitempty" xml:"type,omitempty"`
    

    // 开始时间（查询关键词报告需要的参数）
    
    BeginDate   string `json:"begin_date,omitempty" xml:"begin_date,omitempty"`
    

    // 结束时间（查询关键词报告需要的参数）
    
    EndDate   string `json:"end_date,omitempty" xml:"end_date,omitempty"`
    

    // 配置ksy
    
    SettingKey   string `json:"setting_key,omitempty" xml:"setting_key,omitempty"`
    

    // 配置value
    
    SettingValue   string `json:"setting_value,omitempty" xml:"setting_value,omitempty"`
    

    // 产品id
    
    ProductId   int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
    

}
