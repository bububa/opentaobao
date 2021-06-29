package singletreasure

// ActivityNameInfoDTO 
type ActivityNameInfoDTO struct {
    // 类目
    Category   string `json:"category,omitempty" xml:"category,omitempty"`
    // 活动名称的列表
    NameList   []string `json:"name_list,omitempty" xml:"name_list>string,omitempty"`
}
