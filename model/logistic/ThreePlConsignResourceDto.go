package logistic

// ThreePlConsignResourceDto 
/* model for simplify = false
type ThreePlConsignResourceDto struct {

    // 破损赔付
    
    BrokenCompensatePrice   int64 `json:"broken_compensate_price,omitempty"`
    

    // 首重价格
    
    BasicWeight   int64 `json:"basic_weight,omitempty"`
    

    // 达成时效
    
    DeliveryTime   int64 `json:"delivery_time,omitempty"`
    

    // 续重价格
    
    StepWeight   int64 `json:"step_weight,omitempty"`
    

    // 首重价格
    
    BasicWeightPrice   int64 `json:"basic_weight_price,omitempty"`
    

    // 资源code
    
    ResCode   string `json:"res_code,omitempty"`
    

    // 资源名称
    
    ResName   string `json:"res_name,omitempty"`
    

    // 达成率
    
    AchievingRate   int64 `json:"achieving_rate,omitempty"`
    

    // 续重价格
    
    StepWeightPrice   int64 `json:"step_weight_price,omitempty"`
    

    // 配送资源id
    
    ResId   int64 `json:"res_id,omitempty"`
    

    // 丢失赔付价格
    
    MissingCompensatePrice   int64 `json:"missing_compensate_price,omitempty"`
    

}
*/

// ThreePlConsignResourceDto 
type ThreePlConsignResourceDto struct {

    // 破损赔付
    BrokenCompensatePrice   int64 `json:"broken_compensate_price,omitempty"`

    // 首重价格
    BasicWeight   int64 `json:"basic_weight,omitempty"`

    // 达成时效
    DeliveryTime   int64 `json:"delivery_time,omitempty"`

    // 续重价格
    StepWeight   int64 `json:"step_weight,omitempty"`

    // 首重价格
    BasicWeightPrice   int64 `json:"basic_weight_price,omitempty"`

    // 资源code
    ResCode   string `json:"res_code,omitempty"`

    // 资源名称
    ResName   string `json:"res_name,omitempty"`

    // 达成率
    AchievingRate   int64 `json:"achieving_rate,omitempty"`

    // 续重价格
    StepWeightPrice   int64 `json:"step_weight_price,omitempty"`

    // 配送资源id
    ResId   int64 `json:"res_id,omitempty"`

    // 丢失赔付价格
    MissingCompensatePrice   int64 `json:"missing_compensate_price,omitempty"`

}
