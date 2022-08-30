package logistic

// ReportShippedRequestDto 结构体
type ReportShippedRequestDto struct {
	// report shipped list
	ReportShippedList []ReportShippedDto `json:"report_shipped_list,omitempty" xml:"report_shipped_list>report_shipped_dto,omitempty"`
}
