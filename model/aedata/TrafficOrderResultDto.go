package aedata

// TrafficOrderResultDto 结构体
type TrafficOrderResultDto struct {
	// 订单内容明细
	Orders []Order `json:"orders,omitempty" xml:"orders>order,omitempty"`
	// 当前返回记录的最大游标ID
	MaxQueryIndexId string `json:"max_query_index_id,omitempty" xml:"max_query_index_id,omitempty"`
	// 当前返回记录的最小游标ID
	MinQueryIndexId string `json:"min_query_index_id,omitempty" xml:"min_query_index_id,omitempty"`
	// 当前页记录条数
	CurrentRecordCount int64 `json:"current_record_count,omitempty" xml:"current_record_count,omitempty"`
	// 当前页码
	CurrentPageNo int64 `json:"current_page_no,omitempty" xml:"current_page_no,omitempty"`
	// 总页数
	TotalPageNo int64 `json:"total_page_no,omitempty" xml:"total_page_no,omitempty"`
	// 总记录数
	TotalRecordCount int64 `json:"total_record_count,omitempty" xml:"total_record_count,omitempty"`
}
