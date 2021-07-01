package promotion

// ActivityRelationDetailRequest 结构体
type ActivityRelationDetailRequest struct {
	// 活动状态(VALID  ， DELETE)
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// ISV活动关联权益后获得的关联ID
	RelationId int64 `json:"relation_id,omitempty" xml:"relation_id,omitempty"`
}
