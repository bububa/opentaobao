package tmallnr

// NrtQueryStoreReq 结构体
type NrtQueryStoreReq struct {
	// 活动ID
	ActivityId int64 `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
	// 每页数据条数，最大100条
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 页码
	PageIndex int64 `json:"page_index,omitempty" xml:"page_index,omitempty"`
}
