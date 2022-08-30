package aedropshiper

// TrafficProductResultDto 结构体
type TrafficProductResultDto struct {
	// products
	Products []Integer `json:"products,omitempty" xml:"products>integer,omitempty"`
	// total record count
	TotalRecordCount int64 `json:"total_record_count,omitempty" xml:"total_record_count,omitempty"`
	// current record count
	CurrentRecordCount int64 `json:"current_record_count,omitempty" xml:"current_record_count,omitempty"`
	// total page number
	TotalPageNo int64 `json:"total_page_no,omitempty" xml:"total_page_no,omitempty"`
	// count page number
	CurrentPageNo int64 `json:"current_page_no,omitempty" xml:"current_page_no,omitempty"`
	// is finished
	IsFinished bool `json:"is_finished,omitempty" xml:"is_finished,omitempty"`
}
