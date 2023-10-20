package jae

import (
	"sync"
)

// ParamDto 结构体
type ParamDto struct {
	// 业务参数，json格式字符串
	BizParam string `json:"biz_param,omitempty" xml:"biz_param,omitempty"`
	// 扩展参数
	ExtraParam string `json:"extra_param,omitempty" xml:"extra_param,omitempty"`
	// 区分业务类型及方法
	BizType string `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
}

var poolParamDto = sync.Pool{
	New: func() any {
		return new(ParamDto)
	},
}

// GetParamDto() 从对象池中获取ParamDto
func GetParamDto() *ParamDto {
	return poolParamDto.Get().(*ParamDto)
}

// ReleaseParamDto 释放ParamDto
func ReleaseParamDto(v *ParamDto) {
	v.BizParam = ""
	v.ExtraParam = ""
	v.BizType = ""
	poolParamDto.Put(v)
}
