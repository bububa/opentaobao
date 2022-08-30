package wdk

// DateTimeSliceCollectionDtoList 结构体
type DateTimeSliceCollectionDtoList struct {
	// 时间片
	TimeSliceList []TimeSliceList `json:"time_slice_list,omitempty" xml:"time_slice_list>time_slice_list,omitempty"`
	// 日期
	Date string `json:"date,omitempty" xml:"date,omitempty"`
}
