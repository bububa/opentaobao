package dmp

// DmpSelectTagOptionDto 结构体
type DmpSelectTagOptionDto struct {
	// 标签选项的提交值
	Values string `json:"values,omitempty" xml:"values,omitempty"`
	// 分组id,在标签获取的接口中会返回
	OptionGroupId int64 `json:"option_group_id,omitempty" xml:"option_group_id,omitempty"`
	// 标签ID
	TagId int64 `json:"tag_id,omitempty" xml:"tag_id,omitempty"`
}
