package alsc

import (
	"sync"
)

// CodeBizAttributeDto 结构体
type CodeBizAttributeDto struct {
	// 码值
	CodeValue string `json:"code_value,omitempty" xml:"code_value,omitempty"`
	// 码值绑定业务主体类型
	SubjectType string `json:"subject_type,omitempty" xml:"subject_type,omitempty"`
	// 码值绑定的业务主体ID
	SubjectId string `json:"subject_id,omitempty" xml:"subject_id,omitempty"`
}

var poolCodeBizAttributeDto = sync.Pool{
	New: func() any {
		return new(CodeBizAttributeDto)
	},
}

// GetCodeBizAttributeDto() 从对象池中获取CodeBizAttributeDto
func GetCodeBizAttributeDto() *CodeBizAttributeDto {
	return poolCodeBizAttributeDto.Get().(*CodeBizAttributeDto)
}

// ReleaseCodeBizAttributeDto 释放CodeBizAttributeDto
func ReleaseCodeBizAttributeDto(v *CodeBizAttributeDto) {
	v.CodeValue = ""
	v.SubjectType = ""
	v.SubjectId = ""
	poolCodeBizAttributeDto.Put(v)
}
