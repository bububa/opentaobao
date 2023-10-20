package scbp

import (
	"sync"
)

// TagEstimateOperationDto 结构体
type TagEstimateOperationDto struct {
	// optionValues
	OptionValues []string `json:"option_values,omitempty" xml:"option_values>string,omitempty"`
}

var poolTagEstimateOperationDto = sync.Pool{
	New: func() any {
		return new(TagEstimateOperationDto)
	},
}

// GetTagEstimateOperationDto() 从对象池中获取TagEstimateOperationDto
func GetTagEstimateOperationDto() *TagEstimateOperationDto {
	return poolTagEstimateOperationDto.Get().(*TagEstimateOperationDto)
}

// ReleaseTagEstimateOperationDto 释放TagEstimateOperationDto
func ReleaseTagEstimateOperationDto(v *TagEstimateOperationDto) {
	v.OptionValues = v.OptionValues[:0]
	poolTagEstimateOperationDto.Put(v)
}
