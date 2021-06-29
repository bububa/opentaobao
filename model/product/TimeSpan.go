package product

// TimeSpan 
type TimeSpan struct {
    // 变更开始时间
    Start   string `json:"start,omitempty" xml:"start,omitempty"`
    // 变更结束时间
    End   string `json:"end,omitempty" xml:"end,omitempty"`
    // 变更时间段内的具体有变更的日期掩码
    ChangeDateMask   string `json:"change_date_mask,omitempty" xml:"change_date_mask,omitempty"`
}
