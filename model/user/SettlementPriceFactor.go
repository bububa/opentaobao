package user

// SettlementPriceFactor 
type SettlementPriceFactor struct {

    // 计价因子属性
    Name   string `json:"name,omitempty"`

    // 计价因子实际值
    Value   int64 `json:"value,omitempty"`

    // 计价说明
    Desc   string `json:"desc,omitempty"`

}
