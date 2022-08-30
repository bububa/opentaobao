package alitripmerchant

// UserRiskRankVo 结构体
type UserRiskRankVo struct {
	// 请求id
	RequestUnionId int64 `json:"request_union_id,omitempty" xml:"request_union_id,omitempty"`
	// 风险等级
	RistRank int64 `json:"rist_rank,omitempty" xml:"rist_rank,omitempty"`
}
