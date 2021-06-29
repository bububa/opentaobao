package logistic

// AlibabaEleFengniaoChainstoreRangesResult 
type AlibabaEleFengniaoChainstoreRangesResult struct {
    // 到达圈标识
    RangeId   int64 `json:"range_id,omitempty" xml:"range_id,omitempty"`
    // ranges
    Ranges   []Range `json:"ranges,omitempty" xml:"ranges>range,omitempty"`
    // 配送圈类型 1-日间 2- 晚上  3-全天
    RangeType   int64 `json:"range_type,omitempty" xml:"range_type,omitempty"`
}
