package mydata

// DateRange 
type DateRange struct {
    // 数据周期结束日期（含）
    EndDate   string `json:"end_date,omitempty" xml:"end_date,omitempty"`
    // 数据周期开始日期（含）
    StartDate   string `json:"start_date,omitempty" xml:"start_date,omitempty"`
}
