package tmic

import (
	"sync"
)

// SubQuestionItemBo 结构体
type SubQuestionItemBo struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 随机组号
	RandomGroupNumber int64 `json:"random_group_number,omitempty" xml:"random_group_number,omitempty"`
}

var poolSubQuestionItemBo = sync.Pool{
	New: func() any {
		return new(SubQuestionItemBo)
	},
}

// GetSubQuestionItemBo() 从对象池中获取SubQuestionItemBo
func GetSubQuestionItemBo() *SubQuestionItemBo {
	return poolSubQuestionItemBo.Get().(*SubQuestionItemBo)
}

// ReleaseSubQuestionItemBo 释放SubQuestionItemBo
func ReleaseSubQuestionItemBo(v *SubQuestionItemBo) {
	v.Code = ""
	v.Description = ""
	v.RandomGroupNumber = 0
	poolSubQuestionItemBo.Put(v)
}
