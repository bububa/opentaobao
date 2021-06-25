package scbp

// TargetTagOperationDto 
type TargetTagOperationDto struct {

    // crowd or region
    Scope   []String `json:"scope,omitempty"`

    // 出价类型：0=出价, 1=溢价，2=过滤, 3=召回
    PriceMode   int64 `json:"price_mode,omitempty"`

}
