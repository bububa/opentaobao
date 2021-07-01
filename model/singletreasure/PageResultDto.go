package singletreasure

// PageResultDto 结构体
type PageResultDto struct {
	// 系统信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 总数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 查询结果
	DataList []ItemDetailInfo `json:"data_list,omitempty" xml:"data_list>item_detail_info,omitempty"`
	// 页码
	PageNumber int64 `json:"page_number,omitempty" xml:"page_number,omitempty"`
	// 系统编码
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
	// 系统执行成功与否
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 返回数量个数
	Size int64 `json:"size,omitempty" xml:"size,omitempty"`
}
