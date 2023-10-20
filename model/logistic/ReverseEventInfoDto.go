package logistic

import (
	"sync"
)

// ReverseEventInfoDto 结构体
type ReverseEventInfoDto struct {
	// 销退单ID
	Id string `json:"id,omitempty" xml:"id,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 扩展字段
	Extra string `json:"extra,omitempty" xml:"extra,omitempty"`
	// 值
	Value string `json:"value,omitempty" xml:"value,omitempty"`
	// 类型(1=销退单状态变更)
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
}

var poolReverseEventInfoDto = sync.Pool{
	New: func() any {
		return new(ReverseEventInfoDto)
	},
}

// GetReverseEventInfoDto() 从对象池中获取ReverseEventInfoDto
func GetReverseEventInfoDto() *ReverseEventInfoDto {
	return poolReverseEventInfoDto.Get().(*ReverseEventInfoDto)
}

// ReleaseReverseEventInfoDto 释放ReverseEventInfoDto
func ReleaseReverseEventInfoDto(v *ReverseEventInfoDto) {
	v.Id = ""
	v.Remark = ""
	v.Extra = ""
	v.Value = ""
	v.Type = 0
	poolReverseEventInfoDto.Put(v)
}
