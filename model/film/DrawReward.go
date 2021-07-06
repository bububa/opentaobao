package film

// DrawReward 结构体
type DrawReward struct {
	// 有效期时间
	GmtExpire string `json:"gmt_expire,omitempty" xml:"gmt_expire,omitempty"`
	// gmtModified
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 奖品描述
	RewardDesc string `json:"reward_desc,omitempty" xml:"reward_desc,omitempty"`
	// 券码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 奖品数量
	RewardCount int64 `json:"reward_count,omitempty" xml:"reward_count,omitempty"`
	// 奖品类型（1：淘票票优惠券）
	RewardType int64 `json:"reward_type,omitempty" xml:"reward_type,omitempty"`
	// 券子类型（1：兑换券 2：代金券）
	CodeType int64 `json:"code_type,omitempty" xml:"code_type,omitempty"`
	// 奖品金额（单位：分）
	CostPrice int64 `json:"cost_price,omitempty" xml:"cost_price,omitempty"`
}
