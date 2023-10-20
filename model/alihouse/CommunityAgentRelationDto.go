package alihouse

import (
	"sync"
)

// CommunityAgentRelationDto 结构体
type CommunityAgentRelationDto struct {
	// 外部经纪人ID
	OuterConsultantId string `json:"outer_consultant_id,omitempty" xml:"outer_consultant_id,omitempty"`
	// 外部服务者ID
	OuterBrokerId string `json:"outer_broker_id,omitempty" xml:"outer_broker_id,omitempty"`
	// 新外部服务者ID
	NewOuterBrokerId string `json:"new_outer_broker_id,omitempty" xml:"new_outer_broker_id,omitempty"`
	// 类型 1-指定专家 2-系统认定专家
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 是否测试 0-否 1-是
	IsTest int64 `json:"is_test,omitempty" xml:"is_test,omitempty"`
}

var poolCommunityAgentRelationDto = sync.Pool{
	New: func() any {
		return new(CommunityAgentRelationDto)
	},
}

// GetCommunityAgentRelationDto() 从对象池中获取CommunityAgentRelationDto
func GetCommunityAgentRelationDto() *CommunityAgentRelationDto {
	return poolCommunityAgentRelationDto.Get().(*CommunityAgentRelationDto)
}

// ReleaseCommunityAgentRelationDto 释放CommunityAgentRelationDto
func ReleaseCommunityAgentRelationDto(v *CommunityAgentRelationDto) {
	v.OuterConsultantId = ""
	v.OuterBrokerId = ""
	v.NewOuterBrokerId = ""
	v.Type = 0
	v.IsTest = 0
	poolCommunityAgentRelationDto.Put(v)
}
