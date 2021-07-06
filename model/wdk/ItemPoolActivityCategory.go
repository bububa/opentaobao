package wdk

// ItemPoolActivityCategory 结构体
type ItemPoolActivityCategory struct {
	// 类目分组
	CategoryId string `json:"category_id,omitempty" xml:"category_id,omitempty"`
	// 分组id
	LogicGroupNumber int64 `json:"logic_group_number,omitempty" xml:"logic_group_number,omitempty"`
}
