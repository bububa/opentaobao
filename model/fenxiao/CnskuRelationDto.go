package fenxiao

// CnskuRelationDto 结构体
type CnskuRelationDto struct {
	// r_quantity:2 代表数量
	TypeAttrMap string `json:"type_attr_map,omitempty" xml:"type_attr_map,omitempty"`
	// 组合货品type:8
	RelationType int64 `json:"relation_type,omitempty" xml:"relation_type,omitempty"`
	// 货主Id
	TargetUserId int64 `json:"target_user_id,omitempty" xml:"target_user_id,omitempty"`
	// 组合货品子品Id
	TargetItemId int64 `json:"target_item_id,omitempty" xml:"target_item_id,omitempty"`
	// 组合货品主品Id，创建主品后自动填充
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}
