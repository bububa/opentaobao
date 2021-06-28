package flight

// BaggageApiBean 
/* model for simplify = false
type BaggageApiBean struct {

    // 单件行李高度限制，单位厘米
    
    Height   int64 `json:"height,omitempty"`
    

    // 行李重量，当allWeight为是的时候行李重量描述所有件数的总重量
    
    Weight   int64 `json:"weight,omitempty"`
    

    // 单件行李宽度限制，单位厘米
    
    Width   int64 `json:"width,omitempty"`
    

    // 行李总件数，大于等于1的正整数
    
    Pc   int64 `json:"pc,omitempty"`
    

    // 单件行李长度限制，单位厘米
    
    Length   int64 `json:"length,omitempty"`
    

    // 是否所有行李重量
    
    IsAllWeight   bool `json:"is_all_weight,omitempty"`
    

}
*/

// BaggageApiBean 
type BaggageApiBean struct {

    // 单件行李高度限制，单位厘米
    Height   int64 `json:"height,omitempty"`

    // 行李重量，当allWeight为是的时候行李重量描述所有件数的总重量
    Weight   int64 `json:"weight,omitempty"`

    // 单件行李宽度限制，单位厘米
    Width   int64 `json:"width,omitempty"`

    // 行李总件数，大于等于1的正整数
    Pc   int64 `json:"pc,omitempty"`

    // 单件行李长度限制，单位厘米
    Length   int64 `json:"length,omitempty"`

    // 是否所有行李重量
    IsAllWeight   bool `json:"is_all_weight,omitempty"`

}
