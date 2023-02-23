package wdk

// PromiseTimeSlice 结构体
type PromiseTimeSlice struct {
	// 时间片小时维度列表, 单个时间片格式为09:00-09:30
	SliceList []string `json:"slice_list,omitempty" xml:"slice_list>string,omitempty"`
	// 时间片日期 YYYY-MM-DD
	SliceDate string `json:"slice_date,omitempty" xml:"slice_date,omitempty"`
}
