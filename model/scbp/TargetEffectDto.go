package scbp

// TargetEffectDto 
/* model for simplify = false
type TargetEffectDto struct {

    // 标题
    
    Subject   string `json:"subject,omitempty"`
    

    // 选项值
    
    OptionValue   string `json:"option_value,omitempty"`
    

    // 13:地域标签 14：人群标签
    
    TagRefType   int64 `json:"tag_ref_type,omitempty"`
    

    // 日期(yyyy-MM-dd)
    
    StatDate   string `json:"stat_date,omitempty"`
    

    // 曝光
    
    Impr   int64 `json:"impr,omitempty"`
    

    // 点击
    
    Click   int64 `json:"click,omitempty"`
    

    // 消耗
    
    Cost   int64 `json:"cost,omitempty"`
    

    // 推广时长
    
    OnlineMin   int64 `json:"online_min,omitempty"`
    

}
*/

// TargetEffectDto 
type TargetEffectDto struct {

    // 标题
    Subject   string `json:"subject,omitempty"`

    // 选项值
    OptionValue   string `json:"option_value,omitempty"`

    // 13:地域标签 14：人群标签
    TagRefType   int64 `json:"tag_ref_type,omitempty"`

    // 日期(yyyy-MM-dd)
    StatDate   string `json:"stat_date,omitempty"`

    // 曝光
    Impr   int64 `json:"impr,omitempty"`

    // 点击
    Click   int64 `json:"click,omitempty"`

    // 消耗
    Cost   int64 `json:"cost,omitempty"`

    // 推广时长
    OnlineMin   int64 `json:"online_min,omitempty"`

}
