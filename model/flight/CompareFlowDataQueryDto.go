package flight

import (
	"sync"
)

// CompareFlowDataQueryDto 结构体
type CompareFlowDataQueryDto struct {
	// 店铺id
	AgentId int64 `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
}

var poolCompareFlowDataQueryDto = sync.Pool{
	New: func() any {
		return new(CompareFlowDataQueryDto)
	},
}

// GetCompareFlowDataQueryDto() 从对象池中获取CompareFlowDataQueryDto
func GetCompareFlowDataQueryDto() *CompareFlowDataQueryDto {
	return poolCompareFlowDataQueryDto.Get().(*CompareFlowDataQueryDto)
}

// ReleaseCompareFlowDataQueryDto 释放CompareFlowDataQueryDto
func ReleaseCompareFlowDataQueryDto(v *CompareFlowDataQueryDto) {
	v.AgentId = 0
	poolCompareFlowDataQueryDto.Put(v)
}
