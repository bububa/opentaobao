package scbp

// AdKeywordEffectDto 
/* model for simplify = false
type AdKeywordEffectDto struct {

    // 关键词
    
    Keyword   string `json:"keyword,omitempty"`
    

    // 曝光
    
    ImpressionCnt   string `json:"impression_cnt,omitempty"`
    

    // 平均点击花费
    
    ClickCostAvg   string `json:"click_cost_avg,omitempty"`
    

    // 单位小时，保留一位小数，例如13.5表示13.5小时
    
    OnlineTime   string `json:"online_time,omitempty"`
    

    // 点击量
    
    ClickCnt   string `json:"click_cnt,omitempty"`
    

    // 单位元，保留两位小数, 例如3.75表示3.75元
    
    Cost   string `json:"cost,omitempty"`
    

    // 百分比，保留两位小数，例如3.75表示3.75%
    
    Ctr   string `json:"ctr,omitempty"`
    

}
*/

// AdKeywordEffectDto 
type AdKeywordEffectDto struct {

    // 关键词
    Keyword   string `json:"keyword,omitempty"`

    // 曝光
    ImpressionCnt   string `json:"impression_cnt,omitempty"`

    // 平均点击花费
    ClickCostAvg   string `json:"click_cost_avg,omitempty"`

    // 单位小时，保留一位小数，例如13.5表示13.5小时
    OnlineTime   string `json:"online_time,omitempty"`

    // 点击量
    ClickCnt   string `json:"click_cnt,omitempty"`

    // 单位元，保留两位小数, 例如3.75表示3.75元
    Cost   string `json:"cost,omitempty"`

    // 百分比，保留两位小数，例如3.75表示3.75%
    Ctr   string `json:"ctr,omitempty"`

}
