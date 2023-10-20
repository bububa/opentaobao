package film

import (
	"sync"
)

// LotteryPerformanceRewardDto 结构体
type LotteryPerformanceRewardDto struct {
	// 奖品主标题
	RewardTitle string `json:"reward_title,omitempty" xml:"reward_title,omitempty"`
	// 奖品副标题
	RewardSubTitle string `json:"reward_sub_title,omitempty" xml:"reward_sub_title,omitempty"`
	// 奖品单位
	RewardUnit string `json:"reward_unit,omitempty" xml:"reward_unit,omitempty"`
	// 奖品过期时间(格式: yyyy-MM-dd HH:mm:ss)
	GmtExpire string `json:"gmt_expire,omitempty" xml:"gmt_expire,omitempty"`
	// 拓展参数
	RewardExt string `json:"reward_ext,omitempty" xml:"reward_ext,omitempty"`
	// 奖品类型（1：淘票票优惠券；2：大麦优惠券；14：会员权益）
	RewardType int64 `json:"reward_type,omitempty" xml:"reward_type,omitempty"`
	// 优惠券子类型（1：兑换券 2：代金券）会员权益子类型（5：会员积分；7：普通会员；8：黑钻会员）
	RewardSubType int64 `json:"reward_sub_type,omitempty" xml:"reward_sub_type,omitempty"`
	// 奖品数量
	RewardCount int64 `json:"reward_count,omitempty" xml:"reward_count,omitempty"`
	// 奖品面额(单位分)
	CostPrice int64 `json:"cost_price,omitempty" xml:"cost_price,omitempty"`
}

var poolLotteryPerformanceRewardDto = sync.Pool{
	New: func() any {
		return new(LotteryPerformanceRewardDto)
	},
}

// GetLotteryPerformanceRewardDto() 从对象池中获取LotteryPerformanceRewardDto
func GetLotteryPerformanceRewardDto() *LotteryPerformanceRewardDto {
	return poolLotteryPerformanceRewardDto.Get().(*LotteryPerformanceRewardDto)
}

// ReleaseLotteryPerformanceRewardDto 释放LotteryPerformanceRewardDto
func ReleaseLotteryPerformanceRewardDto(v *LotteryPerformanceRewardDto) {
	v.RewardTitle = ""
	v.RewardSubTitle = ""
	v.RewardUnit = ""
	v.GmtExpire = ""
	v.RewardExt = ""
	v.RewardType = 0
	v.RewardSubType = 0
	v.RewardCount = 0
	v.CostPrice = 0
	poolLotteryPerformanceRewardDto.Put(v)
}
