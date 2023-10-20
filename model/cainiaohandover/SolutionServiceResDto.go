package cainiaohandover

import (
	"sync"
)

// SolutionServiceResDto 结构体
type SolutionServiceResDto struct {
	// 解决方案编码
	SolutionCode string `json:"solution_code,omitempty" xml:"solution_code,omitempty"`
	// 优先级
	Priority string `json:"priority,omitempty" xml:"priority,omitempty"`
	// 联系人名称
	ContactName string `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
	// 联系人电话
	ContactTelephone string `json:"contact_telephone,omitempty" xml:"contact_telephone,omitempty"`
	// 工作时间
	WorkTimeTips string `json:"work_time_tips,omitempty" xml:"work_time_tips,omitempty"`
	// 地址对应的divisionId
	Division string `json:"division,omitempty" xml:"division,omitempty"`
	// 地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 资源名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 资源编码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 扩展字段
	Features *Features `json:"features,omitempty" xml:"features,omitempty"`
}

var poolSolutionServiceResDto = sync.Pool{
	New: func() any {
		return new(SolutionServiceResDto)
	},
}

// GetSolutionServiceResDto() 从对象池中获取SolutionServiceResDto
func GetSolutionServiceResDto() *SolutionServiceResDto {
	return poolSolutionServiceResDto.Get().(*SolutionServiceResDto)
}

// ReleaseSolutionServiceResDto 释放SolutionServiceResDto
func ReleaseSolutionServiceResDto(v *SolutionServiceResDto) {
	v.SolutionCode = ""
	v.Priority = ""
	v.ContactName = ""
	v.ContactTelephone = ""
	v.WorkTimeTips = ""
	v.Division = ""
	v.Address = ""
	v.Name = ""
	v.Code = ""
	v.Features = nil
	poolSolutionServiceResDto.Put(v)
}
