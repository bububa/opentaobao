package alitripmerchant

import (
	"sync"
)

// UserRiskRankVo 结构体
type UserRiskRankVo struct {
	// 请求id
	RequestUnionId int64 `json:"request_union_id,omitempty" xml:"request_union_id,omitempty"`
	// 风险等级
	RistRank int64 `json:"rist_rank,omitempty" xml:"rist_rank,omitempty"`
}

var poolUserRiskRankVo = sync.Pool{
	New: func() any {
		return new(UserRiskRankVo)
	},
}

// GetUserRiskRankVo() 从对象池中获取UserRiskRankVo
func GetUserRiskRankVo() *UserRiskRankVo {
	return poolUserRiskRankVo.Get().(*UserRiskRankVo)
}

// ReleaseUserRiskRankVo 释放UserRiskRankVo
func ReleaseUserRiskRankVo(v *UserRiskRankVo) {
	v.RequestUnionId = 0
	v.RistRank = 0
	poolUserRiskRankVo.Put(v)
}
