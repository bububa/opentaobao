package alihouse

// GradeItemDTOList 结构体
type GradeItemDTOList struct {
	// 子项列表
	StoreLevelSubItemDTOList []StoreLevelSubItemDto `json:"store_level_sub_item_d_t_o_list,omitempty" xml:"store_level_sub_item_d_t_o_list>store_level_sub_item_dto,omitempty"`
	// 权益名称
	GradeItemName string `json:"grade_item_name,omitempty" xml:"grade_item_name,omitempty"`
	// 权益编号
	GradeItemCode string `json:"grade_item_code,omitempty" xml:"grade_item_code,omitempty"`
	// 权益要求等分
	Score int64 `json:"score,omitempty" xml:"score,omitempty"`
}
