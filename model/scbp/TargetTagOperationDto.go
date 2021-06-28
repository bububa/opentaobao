package scbp

// TargetTagOperationDto 
/* model for simplify = false
type TargetTagOperationDto struct {

    // crowd or region
    
    Scope  struct {
        String  []string `json:"string,omitempty"`
    } `json:"scope,omitempty"`
    

    // 出价类型：0=出价, 1=溢价，2=过滤, 3=召回
    
    PriceMode   int64 `json:"price_mode,omitempty"`
    

}
*/

// TargetTagOperationDto 
type TargetTagOperationDto struct {

    // crowd or region
    Scope   []string `json:"scope,omitempty"`

    // 出价类型：0=出价, 1=溢价，2=过滤, 3=召回
    PriceMode   int64 `json:"price_mode,omitempty"`

}
