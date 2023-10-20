package cainiaohandover

import (
	"sync"
)

// OpenSolutionDto 结构体
type OpenSolutionDto struct {
	// 时效信息
	TimingList []OpenTimingDto `json:"timing_list,omitempty" xml:"timing_list>open_timing_dto,omitempty"`
	// 费用列表
	FeeList []OpenFeeDto `json:"fee_list,omitempty" xml:"fee_list>open_fee_dto,omitempty"`
	// 解决方案code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 解决方案名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 推荐指数
	RecommendIndex int64 `json:"recommend_index,omitempty" xml:"recommend_index,omitempty"`
}

var poolOpenSolutionDto = sync.Pool{
	New: func() any {
		return new(OpenSolutionDto)
	},
}

// GetOpenSolutionDto() 从对象池中获取OpenSolutionDto
func GetOpenSolutionDto() *OpenSolutionDto {
	return poolOpenSolutionDto.Get().(*OpenSolutionDto)
}

// ReleaseOpenSolutionDto 释放OpenSolutionDto
func ReleaseOpenSolutionDto(v *OpenSolutionDto) {
	v.TimingList = v.TimingList[:0]
	v.FeeList = v.FeeList[:0]
	v.Code = ""
	v.Name = ""
	v.RecommendIndex = 0
	poolOpenSolutionDto.Put(v)
}
