package feedflow

// TimeSpanDto 
type TimeSpanDto struct {
    // 折扣率
    Discount   int64 `json:"discount,omitempty" xml:"discount,omitempty"`
    // 时间
    Time   string `json:"time,omitempty" xml:"time,omitempty"`
}
