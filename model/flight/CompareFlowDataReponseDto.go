package flight

import (
	"sync"
)

// CompareFlowDataReponseDto 结构体
type CompareFlowDataReponseDto struct {
	// 店铺id
	AgentId int64 `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
	// 当日已使用量
	CurrentAmount int64 `json:"current_amount,omitempty" xml:"current_amount,omitempty"`
	// 当日分配给商家的总调用量
	TotalLimit int64 `json:"total_limit,omitempty" xml:"total_limit,omitempty"`
}

var poolCompareFlowDataReponseDto = sync.Pool{
	New: func() any {
		return new(CompareFlowDataReponseDto)
	},
}

// GetCompareFlowDataReponseDto() 从对象池中获取CompareFlowDataReponseDto
func GetCompareFlowDataReponseDto() *CompareFlowDataReponseDto {
	return poolCompareFlowDataReponseDto.Get().(*CompareFlowDataReponseDto)
}

// ReleaseCompareFlowDataReponseDto 释放CompareFlowDataReponseDto
func ReleaseCompareFlowDataReponseDto(v *CompareFlowDataReponseDto) {
	v.AgentId = 0
	v.CurrentAmount = 0
	v.TotalLimit = 0
	poolCompareFlowDataReponseDto.Put(v)
}
