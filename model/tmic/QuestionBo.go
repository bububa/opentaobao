package tmic

import (
	"sync"
)

// QuestionBo 结构体
type QuestionBo struct {
	// 此问题唯一编码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 问题类型，1-单选，2-多选，3-单行文本，4-多行文本
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 问卷描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 问卷提示
	Tip string `json:"tip,omitempty" xml:"tip,omitempty"`
	// 错误提示
	ErrTip string `json:"err_tip,omitempty" xml:"err_tip,omitempty"`
	// 文本控件占位符
	Placeholder string `json:"placeholder,omitempty" xml:"placeholder,omitempty"`
	// 该问题下选项对象
	Options *OptionBo `json:"options,omitempty" xml:"options,omitempty"`
	// 问题排序字段
	OrderNumber int64 `json:"order_number,omitempty" xml:"order_number,omitempty"`
	// 选项最大个数，5表示最多5个选项
	OptionLimit int64 `json:"option_limit,omitempty" xml:"option_limit,omitempty"`
	// optionBO
	OptionBo *OptionBo `json:"option_bo,omitempty" xml:"option_bo,omitempty"`
	// 子问题
	SubQuestionBO *SubQuestionBo `json:"sub_question_b_o,omitempty" xml:"sub_question_b_o,omitempty"`
	// 该题是否必答，true-必答，false-选答
	Required bool `json:"required,omitempty" xml:"required,omitempty"`
}

var poolQuestionBo = sync.Pool{
	New: func() any {
		return new(QuestionBo)
	},
}

// GetQuestionBo() 从对象池中获取QuestionBo
func GetQuestionBo() *QuestionBo {
	return poolQuestionBo.Get().(*QuestionBo)
}

// ReleaseQuestionBo 释放QuestionBo
func ReleaseQuestionBo(v *QuestionBo) {
	v.Code = ""
	v.Type = ""
	v.Description = ""
	v.Tip = ""
	v.ErrTip = ""
	v.Placeholder = ""
	v.Options = nil
	v.OrderNumber = 0
	v.OptionLimit = 0
	v.OptionBo = nil
	v.SubQuestionBO = nil
	v.Required = false
	poolQuestionBo.Put(v)
}
