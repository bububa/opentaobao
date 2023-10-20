package alihealth2

import (
	"sync"
)

// MedicalBaseTopRequestDto 结构体
type MedicalBaseTopRequestDto struct {
	// 评论业务数据
	BizContent string `json:"biz_content,omitempty" xml:"biz_content,omitempty"`
}

var poolMedicalBaseTopRequestDto = sync.Pool{
	New: func() any {
		return new(MedicalBaseTopRequestDto)
	},
}

// GetMedicalBaseTopRequestDto() 从对象池中获取MedicalBaseTopRequestDto
func GetMedicalBaseTopRequestDto() *MedicalBaseTopRequestDto {
	return poolMedicalBaseTopRequestDto.Get().(*MedicalBaseTopRequestDto)
}

// ReleaseMedicalBaseTopRequestDto 释放MedicalBaseTopRequestDto
func ReleaseMedicalBaseTopRequestDto(v *MedicalBaseTopRequestDto) {
	v.BizContent = ""
	poolMedicalBaseTopRequestDto.Put(v)
}
