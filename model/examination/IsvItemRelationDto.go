package examination

// IsvItemRelationDto 结构体
type IsvItemRelationDto struct {
	// 单项id
	IsvItemId string `json:"isv_item_id,omitempty" xml:"isv_item_id,omitempty"`
	// 关联单项id
	RelIsvItemId string `json:"rel_isv_item_id,omitempty" xml:"rel_isv_item_id,omitempty"`
	// 关系：1、互斥 2、重复
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
}
