package scbp

// BasicQuickCampaign 
/* model for simplify = false
type BasicQuickCampaign struct {

    // 出价区间-上限(可以存在1位小数，大于等于3)
    
    MaxPrice   string `json:"max_price,omitempty"`
    

    // 出价区间-下限(可以存在1位小数)
    
    MinPrice   string `json:"min_price,omitempty"`
    

    // 每日预算,不低于50元
    
    Budget   int64 `json:"budget,omitempty"`
    

    // 计划名称,不能大于20字符
    
    Title   string `json:"title,omitempty"`
    

}
*/

// BasicQuickCampaign 
type BasicQuickCampaign struct {

    // 出价区间-上限(可以存在1位小数，大于等于3)
    MaxPrice   string `json:"max_price,omitempty"`

    // 出价区间-下限(可以存在1位小数)
    MinPrice   string `json:"min_price,omitempty"`

    // 每日预算,不低于50元
    Budget   int64 `json:"budget,omitempty"`

    // 计划名称,不能大于20字符
    Title   string `json:"title,omitempty"`

}
