package tmic

// SubQuestionItemBo 结构体
type SubQuestionItemBo struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 随机组号
	RandomGroupNumber int64 `json:"random_group_number,omitempty" xml:"random_group_number,omitempty"`
}
