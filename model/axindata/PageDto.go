package axindata

// PageDto 结构体
type PageDto struct {
	// 标准酒店id列表
	DataList []int64 `json:"data_list,omitempty" xml:"data_list>int64,omitempty"`
	// 记录数
	Count int64 `json:"count,omitempty" xml:"count,omitempty"`
	// 页码
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// 页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 上次返回最大酒店id
	MaxDateId int64 `json:"max_date_id,omitempty" xml:"max_date_id,omitempty"`
}
