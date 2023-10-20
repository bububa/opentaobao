package alihouse

import (
	"sync"
)

// RegionExpertResultDto 结构体
type RegionExpertResultDto struct {
	// 外部经纪人ID
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 异常信息
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolRegionExpertResultDto = sync.Pool{
	New: func() any {
		return new(RegionExpertResultDto)
	},
}

// GetRegionExpertResultDto() 从对象池中获取RegionExpertResultDto
func GetRegionExpertResultDto() *RegionExpertResultDto {
	return poolRegionExpertResultDto.Get().(*RegionExpertResultDto)
}

// ReleaseRegionExpertResultDto 释放RegionExpertResultDto
func ReleaseRegionExpertResultDto(v *RegionExpertResultDto) {
	v.OuterId = ""
	v.Msg = ""
	v.Success = false
	poolRegionExpertResultDto.Put(v)
}
