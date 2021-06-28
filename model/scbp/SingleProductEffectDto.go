package scbp

// SingleProductEffectDto 
/* model for simplify = false
type SingleProductEffectDto struct {

    // 产品title
    
    Subject   string `json:"subject,omitempty"`
    

    // 曝光
    
    ImpressionCnt   string `json:"impression_cnt,omitempty"`
    

    // 日期
    
    StatDate   string `json:"stat_date,omitempty"`
    

    // 产品ID
    
    ProductId   int64 `json:"product_id,omitempty"`
    

    // 平均点击花费
    
    ClickCostAvg   string `json:"click_cost_avg,omitempty"`
    

    // 点击数
    
    ClickCnt   string `json:"click_cnt,omitempty"`
    

    // 单位元，保留两位小数, 例如3.75表示3.75元
    
    Cost   string `json:"cost,omitempty"`
    

    // 百分比，保留两位小数，例如3.75表示3.75%
    
    Ctr   string `json:"ctr,omitempty"`
    

}
*/

// SingleProductEffectDto 
type SingleProductEffectDto struct {

    // 产品title
    Subject   string `json:"subject,omitempty"`

    // 曝光
    ImpressionCnt   string `json:"impression_cnt,omitempty"`

    // 日期
    StatDate   string `json:"stat_date,omitempty"`

    // 产品ID
    ProductId   int64 `json:"product_id,omitempty"`

    // 平均点击花费
    ClickCostAvg   string `json:"click_cost_avg,omitempty"`

    // 点击数
    ClickCnt   string `json:"click_cnt,omitempty"`

    // 单位元，保留两位小数, 例如3.75表示3.75元
    Cost   string `json:"cost,omitempty"`

    // 百分比，保留两位小数，例如3.75表示3.75%
    Ctr   string `json:"ctr,omitempty"`

}
