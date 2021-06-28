package servicecenter

// ArticleItemViewUnit 
/* model for simplify = false
type ArticleItemViewUnit struct {

    // 收费项目code
    
    ItemCode   string `json:"item_code,omitempty"`
    

    // 收费项目名称
    
    ItemName   string `json:"item_name,omitempty"`
    

    // 周期数，如1，3，6，12。对于周期型和周期计量型返回。
    
    CycNum   int64 `json:"cyc_num,omitempty"`
    

    // 1-年，2-月，3-日。对于周期型和周期计量型返回。
    
    CycUnit   int64 `json:"cyc_unit,omitempty"`
    

    // 数量。对于周期计量型返回计量数。
    
    Quantity   int64 `json:"quantity,omitempty"`
    

    // 原价，单位：元
    
    OriginPrice   string `json:"origin_price,omitempty"`
    

    // 优惠，单位：元
    
    PromPrice   string `json:"prom_price,omitempty"`
    

    // 需要支付的价格，单位：元
    
    ActualPrice   string `json:"actual_price,omitempty"`
    

    // 用户是否可以购买
    
    CanSub   bool `json:"can_sub,omitempty"`
    

    // 错误码
    
    ErrorCode   string `json:"error_code,omitempty"`
    

    // 错误文案
    
    ErrorMsg   string `json:"error_msg,omitempty"`
    

}
*/

// ArticleItemViewUnit 
type ArticleItemViewUnit struct {

    // 收费项目code
    ItemCode   string `json:"item_code,omitempty"`

    // 收费项目名称
    ItemName   string `json:"item_name,omitempty"`

    // 周期数，如1，3，6，12。对于周期型和周期计量型返回。
    CycNum   int64 `json:"cyc_num,omitempty"`

    // 1-年，2-月，3-日。对于周期型和周期计量型返回。
    CycUnit   int64 `json:"cyc_unit,omitempty"`

    // 数量。对于周期计量型返回计量数。
    Quantity   int64 `json:"quantity,omitempty"`

    // 原价，单位：元
    OriginPrice   string `json:"origin_price,omitempty"`

    // 优惠，单位：元
    PromPrice   string `json:"prom_price,omitempty"`

    // 需要支付的价格，单位：元
    ActualPrice   string `json:"actual_price,omitempty"`

    // 用户是否可以购买
    CanSub   bool `json:"can_sub,omitempty"`

    // 错误码
    ErrorCode   string `json:"error_code,omitempty"`

    // 错误文案
    ErrorMsg   string `json:"error_msg,omitempty"`

}
