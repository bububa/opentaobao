package store

// StoreRelationSimpleQuery 结构体
type StoreRelationSimpleQuery struct {
	// 业务类型
	BizType int64 `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
	// 外部关联业务id
	OuterId int64 `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 页码
	PageIndex int64 `json:"page_index,omitempty" xml:"page_index,omitempty"`
	// 每页条目数（最大值100）
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 门店id
	StoreId int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 是否获取精确数量，设置为true后搜索性能较差，对准确总数有强诉求的场景再使用。分页时请只在第一页请求时设置为true
	AccurateCount bool `json:"accurate_count,omitempty" xml:"accurate_count,omitempty"`
}
