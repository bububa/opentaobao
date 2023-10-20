package alihouse

// LevelDtolist 结构体
type LevelDtolist struct {
	// 权益列表
	GradeItemDTOList []GradeItemDtolist `json:"grade_item_d_t_o_list,omitempty" xml:"grade_item_d_t_o_list>grade_item_dtolist,omitempty"`
	// 等级编号
	LevelCode string `json:"level_code,omitempty" xml:"level_code,omitempty"`
	// 外部门店ID
	OuterStoreId string `json:"outer_store_id,omitempty" xml:"outer_store_id,omitempty"`
	// 等级名称
	LevelName string `json:"level_name,omitempty" xml:"level_name,omitempty"`
	// 总分
	TotalScore int64 `json:"total_score,omitempty" xml:"total_score,omitempty"`
}
