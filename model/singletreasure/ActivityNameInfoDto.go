package singletreasure

// ActivityNameInfoDto 结构体
type ActivityNameInfoDto struct {
	// 活动名称的列表
	NameList []string `json:"name_list,omitempty" xml:"name_list>string,omitempty"`
	// 类目
	Category string `json:"category,omitempty" xml:"category,omitempty"`
}
