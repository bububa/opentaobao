package seaking

// ReportQueryApiDto 结构体
type ReportQueryApiDto struct {
	// 请求日期
	QueryDate string `json:"query_date,omitempty" xml:"query_date,omitempty"`
	// 业务类型
	Biz string `json:"biz,omitempty" xml:"biz,omitempty"`
	// 每页size
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 页码
	Page int64 `json:"page,omitempty" xml:"page,omitempty"`
}
