package simba

// RptResult 
/* model for simplify = false
type RptResult struct {

    // ctr
    
    Ctr   string `json:"ctr,omitempty"`
    

    // 消耗
    
    Cost   string `json:"cost,omitempty"`
    

    // 词包类型
    
    IsAutomatch   string `json:"is_automatch,omitempty"`
    

    // cpc
    
    Cpc   string `json:"cpc,omitempty"`
    

    // 时段编号
    
    TimePeriod   string `json:"time_period,omitempty"`
    

    // 展现量
    
    Impression   string `json:"impression,omitempty"`
    

    // 点击量
    
    Click   string `json:"click,omitempty"`
    

}
*/

// RptResult 
type RptResult struct {

    // ctr
    Ctr   string `json:"ctr,omitempty"`

    // 消耗
    Cost   string `json:"cost,omitempty"`

    // 词包类型
    IsAutomatch   string `json:"is_automatch,omitempty"`

    // cpc
    Cpc   string `json:"cpc,omitempty"`

    // 时段编号
    TimePeriod   string `json:"time_period,omitempty"`

    // 展现量
    Impression   string `json:"impression,omitempty"`

    // 点击量
    Click   string `json:"click,omitempty"`

}
