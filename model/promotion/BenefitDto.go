package promotion

// BenefitDto 
/* model for simplify = false
type BenefitDto struct {

    // 权益code
    
    Code   string `json:"code,omitempty"`
    

    // 发放结束时间
    
    EndDate   string `json:"end_date,omitempty"`
    

    // 权益描述
    
    Description   string `json:"description,omitempty"`
    

    // 发放量
    
    Bestow   int64 `json:"bestow,omitempty"`
    

    // 权益类型
    
    Type   string `json:"type,omitempty"`
    

    // 发放开始时间
    
    StartDate   string `json:"start_date,omitempty"`
    

    // 总库存
    
    Total   int64 `json:"total,omitempty"`
    

    // 权益标签
    
    LabelCodes   string `json:"label_codes,omitempty"`
    

    // 扩展信息
    
    Feature  *struct {
        Feature  *Feature `json:"feature,omitempty"`
    } `json:"feature,omitempty"`
    

    // 权益名称
    
    Name   string `json:"name,omitempty"`
    

    // 权益状态
    
    Status   string `json:"status,omitempty"`
    

}
*/

// BenefitDto 
type BenefitDto struct {

    // 权益code
    Code   string `json:"code,omitempty"`

    // 发放结束时间
    EndDate   string `json:"end_date,omitempty"`

    // 权益描述
    Description   string `json:"description,omitempty"`

    // 发放量
    Bestow   int64 `json:"bestow,omitempty"`

    // 权益类型
    Type   string `json:"type,omitempty"`

    // 发放开始时间
    StartDate   string `json:"start_date,omitempty"`

    // 总库存
    Total   int64 `json:"total,omitempty"`

    // 权益标签
    LabelCodes   string `json:"label_codes,omitempty"`

    // 扩展信息
    Feature   *Feature `json:"feature,omitempty"`

    // 权益名称
    Name   string `json:"name,omitempty"`

    // 权益状态
    Status   string `json:"status,omitempty"`

}
