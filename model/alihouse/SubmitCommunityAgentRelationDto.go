package alihouse

// SubmitCommunityAgentRelationDto 结构体
type SubmitCommunityAgentRelationDto struct {
	// 专家列表
	OutConsulatantInfos []CommunityAgentRelationDto `json:"out_consulatant_infos,omitempty" xml:"out_consulatant_infos>community_agent_relation_dto,omitempty"`
	// 外部小区ID
	OuterCommunityId string `json:"outer_community_id,omitempty" xml:"outer_community_id,omitempty"`
	// 业务域
	BusinessType int64 `json:"business_type,omitempty" xml:"business_type,omitempty"`
}
