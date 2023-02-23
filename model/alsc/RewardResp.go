package alsc

// RewardResp 结构体
type RewardResp struct {
	// 任务列表
	RList []RewardInfoDto `json:"r_list,omitempty" xml:"r_list>reward_info_dto,omitempty"`
}
