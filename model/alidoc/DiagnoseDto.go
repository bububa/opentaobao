package alidoc

import (
	"sync"
)

// DiagnoseDto 结构体
type DiagnoseDto struct {
	// icdCode
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// icdCode名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}

var poolDiagnoseDto = sync.Pool{
	New: func() any {
		return new(DiagnoseDto)
	},
}

// GetDiagnoseDto() 从对象池中获取DiagnoseDto
func GetDiagnoseDto() *DiagnoseDto {
	return poolDiagnoseDto.Get().(*DiagnoseDto)
}

// ReleaseDiagnoseDto 释放DiagnoseDto
func ReleaseDiagnoseDto(v *DiagnoseDto) {
	v.Code = ""
	v.Name = ""
	poolDiagnoseDto.Put(v)
}
