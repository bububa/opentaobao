package wlbimports

// DeliveryPrice 
type DeliveryPrice struct {

    // 首重
    BasicWeight   int64 `json:"basic_weight,omitempty"`

    // 首重价格
    BasicWeightPrice   int64 `json:"basic_weight_price,omitempty"`

    // 续重
    StepWeight   int64 `json:"step_weight,omitempty"`

    // 续重价格
    StepWeightPrice   int64 `json:"step_weight_price,omitempty"`

}
