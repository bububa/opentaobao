package alihouse

// StoreLevelSubItemDto 结构体
type StoreLevelSubItemDto struct {
	// 子项名称
	SubItemName string `json:"sub_item_name,omitempty" xml:"sub_item_name,omitempty"`
	// 子项指标值
	SubItemValue string `json:"sub_item_value,omitempty" xml:"sub_item_value,omitempty"`
	// 子项指标目标值
	SubItemTargetValue string `json:"sub_item_target_value,omitempty" xml:"sub_item_target_value,omitempty"`
	// score
	Score string `json:"score,omitempty" xml:"score,omitempty"`
	// 子项目标分
	TargetScore int64 `json:"target_score,omitempty" xml:"target_score,omitempty"`
	// 排序键
	Sort int64 `json:"sort,omitempty" xml:"sort,omitempty"`
}
