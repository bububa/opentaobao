package tmallnr

import (
	"sync"
)

// NrZqsPlanDetailInfoDto 结构体
type NrZqsPlanDetailInfoDto struct {
	// 计划配送时间
	PlanDate string `json:"plan_date,omitempty" xml:"plan_date,omitempty"`
	// 配送期号从1开始，一直到N
	SequenceNo int64 `json:"sequence_no,omitempty" xml:"sequence_no,omitempty"`
}

var poolNrZqsPlanDetailInfoDto = sync.Pool{
	New: func() any {
		return new(NrZqsPlanDetailInfoDto)
	},
}

// GetNrZqsPlanDetailInfoDto() 从对象池中获取NrZqsPlanDetailInfoDto
func GetNrZqsPlanDetailInfoDto() *NrZqsPlanDetailInfoDto {
	return poolNrZqsPlanDetailInfoDto.Get().(*NrZqsPlanDetailInfoDto)
}

// ReleaseNrZqsPlanDetailInfoDto 释放NrZqsPlanDetailInfoDto
func ReleaseNrZqsPlanDetailInfoDto(v *NrZqsPlanDetailInfoDto) {
	v.PlanDate = ""
	v.SequenceNo = 0
	poolNrZqsPlanDetailInfoDto.Put(v)
}
