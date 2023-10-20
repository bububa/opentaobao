package axindata

import (
	"sync"
)

// FscProjectMapDto 结构体
type FscProjectMapDto struct {
	// 供应商（外部）团期编码
	ProjectCode string `json:"project_code,omitempty" xml:"project_code,omitempty"`
	// 供销平台团期编码
	FscProjectId string `json:"fsc_project_id,omitempty" xml:"fsc_project_id,omitempty"`
}

var poolFscProjectMapDto = sync.Pool{
	New: func() any {
		return new(FscProjectMapDto)
	},
}

// GetFscProjectMapDto() 从对象池中获取FscProjectMapDto
func GetFscProjectMapDto() *FscProjectMapDto {
	return poolFscProjectMapDto.Get().(*FscProjectMapDto)
}

// ReleaseFscProjectMapDto 释放FscProjectMapDto
func ReleaseFscProjectMapDto(v *FscProjectMapDto) {
	v.ProjectCode = ""
	v.FscProjectId = ""
	poolFscProjectMapDto.Put(v)
}
