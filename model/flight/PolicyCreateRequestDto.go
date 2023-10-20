package flight

import (
	"sync"
)

// PolicyCreateRequestDto 结构体
type PolicyCreateRequestDto struct {
	// 全量上传时，删除政策查询参数
	DeletePolicy []PolicyQueryParamDto `json:"delete_policy,omitempty" xml:"delete_policy>policy_query_param_dto,omitempty"`
	// 需要导入的政策列表
	PolicyList []PolicyDto `json:"policy_list,omitempty" xml:"policy_list>policy_dto,omitempty"`
	// 操作类型（增量/全量），ADD上传不会删除历史政策，FULL上传会删除所有历史政策，请谨慎使用
	ExecType string `json:"exec_type,omitempty" xml:"exec_type,omitempty"`
	// 政策类型：P，普通政策；1，特殊政策；T，规则政策；G，派单政策
	PolicyType string `json:"policy_type,omitempty" xml:"policy_type,omitempty"`
	// 商家上传时间
	AgentSendTime string `json:"agent_send_time,omitempty" xml:"agent_send_time,omitempty"`
	// 店铺id
	AgentId int64 `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
	// 行程类型：0，单程；1，往返
	TripType int64 `json:"trip_type,omitempty" xml:"trip_type,omitempty"`
}

var poolPolicyCreateRequestDto = sync.Pool{
	New: func() any {
		return new(PolicyCreateRequestDto)
	},
}

// GetPolicyCreateRequestDto() 从对象池中获取PolicyCreateRequestDto
func GetPolicyCreateRequestDto() *PolicyCreateRequestDto {
	return poolPolicyCreateRequestDto.Get().(*PolicyCreateRequestDto)
}

// ReleasePolicyCreateRequestDto 释放PolicyCreateRequestDto
func ReleasePolicyCreateRequestDto(v *PolicyCreateRequestDto) {
	v.DeletePolicy = v.DeletePolicy[:0]
	v.PolicyList = v.PolicyList[:0]
	v.ExecType = ""
	v.PolicyType = ""
	v.AgentSendTime = ""
	v.AgentId = 0
	v.TripType = 0
	poolPolicyCreateRequestDto.Put(v)
}
