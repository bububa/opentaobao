package store

// StoreRelationSimpleDo 结构体
type StoreRelationSimpleDo struct {
	// 业务类型
	BizType int64 `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
	// 门店名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 业务id
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 关系id
	RelationId int64 `json:"relation_id,omitempty" xml:"relation_id,omitempty"`
	// 业务关系状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 门店id
	StoreId int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
}
