package scbp

// ProductEffectDto 
type ProductEffectDto struct {

    // 产品id
    
    ProductId   int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
    

    // 产品名称
    
    ProductName   string `json:"product_name,omitempty" xml:"product_name,omitempty"`
    

    // 图片url
    
    ImgUrl   string `json:"img_url,omitempty" xml:"img_url,omitempty"`
    

    // 日期(yyyy-MM-dd)
    
    StatDate   string `json:"stat_date,omitempty" xml:"stat_date,omitempty"`
    

    // 曝光
    
    Impr   int64 `json:"impr,omitempty" xml:"impr,omitempty"`
    

    // 点击
    
    Click   int64 `json:"click,omitempty" xml:"click,omitempty"`
    

    // 消耗
    
    Cost   int64 `json:"cost,omitempty" xml:"cost,omitempty"`
    

    // 推广时长
    
    OnlineMin   int64 `json:"online_min,omitempty" xml:"online_min,omitempty"`
    

    // title
    
    Subject   string `json:"subject,omitempty" xml:"subject,omitempty"`
    

    // 曝光
    
    ImpressionCnt   string `json:"impression_cnt,omitempty" xml:"impression_cnt,omitempty"`
    

    // 平均点击花费
    
    ClickCostAvg   string `json:"click_cost_avg,omitempty" xml:"click_cost_avg,omitempty"`
    

    // 点击量
    
    ClickCnt   string `json:"click_cnt,omitempty" xml:"click_cnt,omitempty"`
    

    // 百分比，保留两位小数，例如3.75表示3.75%
    
    Ctr   string `json:"ctr,omitempty" xml:"ctr,omitempty"`
    

}
