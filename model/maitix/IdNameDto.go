package maitix

import (
	"sync"
)

// IdNameDto 结构体
type IdNameDto struct {
	// 国家名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 国家id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolIdNameDto = sync.Pool{
	New: func() any {
		return new(IdNameDto)
	},
}

// GetIdNameDto() 从对象池中获取IdNameDto
func GetIdNameDto() *IdNameDto {
	return poolIdNameDto.Get().(*IdNameDto)
}

// ReleaseIdNameDto 释放IdNameDto
func ReleaseIdNameDto(v *IdNameDto) {
	v.Name = ""
	v.Id = 0
	poolIdNameDto.Put(v)
}
