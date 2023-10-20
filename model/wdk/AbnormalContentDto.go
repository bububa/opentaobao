package wdk

import (
	"sync"
)

// AbnormalContentDto 结构体
type AbnormalContentDto struct {
	// 操作描述
	OperateDesc string `json:"operate_desc,omitempty" xml:"operate_desc,omitempty"`
	// 处理结果
	OperateResult string `json:"operate_result,omitempty" xml:"operate_result,omitempty"`
	// 处理数量
	AbnormalProcessQuantity string `json:"abnormal_process_quantity,omitempty" xml:"abnormal_process_quantity,omitempty"`
	// 上报数量
	AbnormalQuantity string `json:"abnormal_quantity,omitempty" xml:"abnormal_quantity,omitempty"`
}

var poolAbnormalContentDto = sync.Pool{
	New: func() any {
		return new(AbnormalContentDto)
	},
}

// GetAbnormalContentDto() 从对象池中获取AbnormalContentDto
func GetAbnormalContentDto() *AbnormalContentDto {
	return poolAbnormalContentDto.Get().(*AbnormalContentDto)
}

// ReleaseAbnormalContentDto 释放AbnormalContentDto
func ReleaseAbnormalContentDto(v *AbnormalContentDto) {
	v.OperateDesc = ""
	v.OperateResult = ""
	v.AbnormalProcessQuantity = ""
	v.AbnormalQuantity = ""
	poolAbnormalContentDto.Put(v)
}
