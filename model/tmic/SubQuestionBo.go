package tmic

import (
	"sync"
)

// SubQuestionBo 结构体
type SubQuestionBo struct {
	// 子问题列表
	SubQuestionItemBOList []SubQuestionItemBo `json:"sub_question_item_b_o_list,omitempty" xml:"sub_question_item_b_o_list>sub_question_item_bo,omitempty"`
}

var poolSubQuestionBo = sync.Pool{
	New: func() any {
		return new(SubQuestionBo)
	},
}

// GetSubQuestionBo() 从对象池中获取SubQuestionBo
func GetSubQuestionBo() *SubQuestionBo {
	return poolSubQuestionBo.Get().(*SubQuestionBo)
}

// ReleaseSubQuestionBo 释放SubQuestionBo
func ReleaseSubQuestionBo(v *SubQuestionBo) {
	v.SubQuestionItemBOList = v.SubQuestionItemBOList[:0]
	poolSubQuestionBo.Put(v)
}
