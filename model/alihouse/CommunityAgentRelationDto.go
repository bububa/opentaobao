package alihouse

// CommunityAgentRelationDto 结构体
type CommunityAgentRelationDto struct {
	// 外部经纪人ID
	OuterConsultantId string `json:"outer_consultant_id,omitempty" xml:"outer_consultant_id,omitempty"`
	// 类型 1-指定专家 2-系统认定专家
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 是否测试 0-否 1-是
	IsTest int64 `json:"is_test,omitempty" xml:"is_test,omitempty"`
}
