package tmic

// SubQuestionBo 结构体
type SubQuestionBo struct {
	// 子问题列表
	SubQuestionItemBOList []SubQuestionItemBo `json:"sub_question_item_b_o_list,omitempty" xml:"sub_question_item_b_o_list>sub_question_item_bo,omitempty"`
}
