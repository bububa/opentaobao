package flight

import (
	"sync"
)

// PolicyTaskQueryDto 结构体
type PolicyTaskQueryDto struct {
	// 店铺id
	AgentId int64 `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
	// 任务id
	TaskId int64 `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

var poolPolicyTaskQueryDto = sync.Pool{
	New: func() any {
		return new(PolicyTaskQueryDto)
	},
}

// GetPolicyTaskQueryDto() 从对象池中获取PolicyTaskQueryDto
func GetPolicyTaskQueryDto() *PolicyTaskQueryDto {
	return poolPolicyTaskQueryDto.Get().(*PolicyTaskQueryDto)
}

// ReleasePolicyTaskQueryDto 释放PolicyTaskQueryDto
func ReleasePolicyTaskQueryDto(v *PolicyTaskQueryDto) {
	v.AgentId = 0
	v.TaskId = 0
	poolPolicyTaskQueryDto.Put(v)
}
