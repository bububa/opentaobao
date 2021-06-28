package logistic

// AlibabaEleFengniaoChainstoreRangesResult 
/* model for simplify = false
type AlibabaEleFengniaoChainstoreRangesResult struct {

    // 到达圈标识
    
    RangeId   int64 `json:"range_id,omitempty"`
    

    // ranges
    
    Ranges  struct {
        Range  []Range `json:"range,omitempty"`
    } `json:"ranges,omitempty"`
    

    // 配送圈类型 1-日间 2- 晚上  3-全天
    
    RangeType   int64 `json:"range_type,omitempty"`
    

}
*/

// AlibabaEleFengniaoChainstoreRangesResult 
type AlibabaEleFengniaoChainstoreRangesResult struct {

    // 到达圈标识
    RangeId   int64 `json:"range_id,omitempty"`

    // ranges
    Ranges   []Range `json:"ranges,omitempty"`

    // 配送圈类型 1-日间 2- 晚上  3-全天
    RangeType   int64 `json:"range_type,omitempty"`

}
