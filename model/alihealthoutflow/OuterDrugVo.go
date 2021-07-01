package alihealthoutflow

// OuterDrugVo 结构体
type OuterDrugVo struct {
	// 医院id
	HosId string `json:"hos_id,omitempty" xml:"hos_id,omitempty"`
	// 分页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 页码
	PageNum int64 `json:"page_num,omitempty" xml:"page_num,omitempty"`
	// 最后查询时间
	LastQueryTime string `json:"last_query_time,omitempty" xml:"last_query_time,omitempty"`
}
