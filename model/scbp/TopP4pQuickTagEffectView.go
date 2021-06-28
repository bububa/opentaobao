package scbp

// TopP4pQuickTagEffectView 
/* model for simplify = false
type TopP4pQuickTagEffectView struct {

    // 标签id(潜在访问偏好和潜在采购意向返回的是类目id，店铺老客和优选人群返回的是字符串)
    
    TagId   string `json:"tag_id,omitempty"`
    

    // 标签名字
    
    TagName   string `json:"tag_name,omitempty"`
    

    // 平均点击出价
    
    Cpc   string `json:"cpc,omitempty"`
    

    // 点击率
    
    Ctr   string `json:"ctr,omitempty"`
    

    // 消耗金额，单位元
    
    Cost   string `json:"cost,omitempty"`
    

    // 点击数
    
    ClickCnt   string `json:"click_cnt,omitempty"`
    

    // 曝光量
    
    ImpressionCnt   string `json:"impression_cnt,omitempty"`
    

}
*/

// TopP4pQuickTagEffectView 
type TopP4pQuickTagEffectView struct {

    // 标签id(潜在访问偏好和潜在采购意向返回的是类目id，店铺老客和优选人群返回的是字符串)
    TagId   string `json:"tag_id,omitempty"`

    // 标签名字
    TagName   string `json:"tag_name,omitempty"`

    // 平均点击出价
    Cpc   string `json:"cpc,omitempty"`

    // 点击率
    Ctr   string `json:"ctr,omitempty"`

    // 消耗金额，单位元
    Cost   string `json:"cost,omitempty"`

    // 点击数
    ClickCnt   string `json:"click_cnt,omitempty"`

    // 曝光量
    ImpressionCnt   string `json:"impression_cnt,omitempty"`

}
