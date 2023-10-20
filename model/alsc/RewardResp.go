package alsc

import (
	"sync"
)

// RewardResp 结构体
type RewardResp struct {
	// 任务列表
	RList []RewardInfoDto `json:"r_list,omitempty" xml:"r_list>reward_info_dto,omitempty"`
}

var poolRewardResp = sync.Pool{
	New: func() any {
		return new(RewardResp)
	},
}

// GetRewardResp() 从对象池中获取RewardResp
func GetRewardResp() *RewardResp {
	return poolRewardResp.Get().(*RewardResp)
}

// ReleaseRewardResp 释放RewardResp
func ReleaseRewardResp(v *RewardResp) {
	v.RList = v.RList[:0]
	poolRewardResp.Put(v)
}
