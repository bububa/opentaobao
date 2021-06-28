package simba

// PriceSuggestionDto 
/* model for simplify = false
type PriceSuggestionDto struct {

    // 关键词id
    
    Bidwordid   string `json:"bidwordid,omitempty"`
    

    // 关键词原词
    
    Word   string `json:"word,omitempty"`
    

    // 状态码
    
    Stat   string `json:"stat,omitempty"`
    

    // 昨日信息
    
    YesterdayInfo  *struct {
        YesterdayInfo  *YesterdayInfo `json:"yesterday_info,omitempty"`
    } `json:"yesterday_info,omitempty"`
    

    // 出价建议
    
    GuidancePrice  *struct {
        GuidancePrice  *GuidancePrice `json:"guidance_price,omitempty"`
    } `json:"guidance_price,omitempty"`
    

}
*/

// PriceSuggestionDto 
type PriceSuggestionDto struct {

    // 关键词id
    Bidwordid   string `json:"bidwordid,omitempty"`

    // 关键词原词
    Word   string `json:"word,omitempty"`

    // 状态码
    Stat   string `json:"stat,omitempty"`

    // 昨日信息
    YesterdayInfo   *YesterdayInfo `json:"yesterday_info,omitempty"`

    // 出价建议
    GuidancePrice   *GuidancePrice `json:"guidance_price,omitempty"`

}
