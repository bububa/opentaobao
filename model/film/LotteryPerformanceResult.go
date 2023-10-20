package film

import (
	"sync"
)

// LotteryPerformanceResult 结构体
type LotteryPerformanceResult struct {
	// 奖品列表
	Rewards []LotteryPerformanceRewardDto `json:"rewards,omitempty" xml:"rewards>lottery_performance_reward_dto,omitempty"`
	// 结果状态码（1：全部成功；2：全部失败；3：部分成功需要重试；4：部分成功无需重试(需人工介入处理)）
	ResultCode int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
}

var poolLotteryPerformanceResult = sync.Pool{
	New: func() any {
		return new(LotteryPerformanceResult)
	},
}

// GetLotteryPerformanceResult() 从对象池中获取LotteryPerformanceResult
func GetLotteryPerformanceResult() *LotteryPerformanceResult {
	return poolLotteryPerformanceResult.Get().(*LotteryPerformanceResult)
}

// ReleaseLotteryPerformanceResult 释放LotteryPerformanceResult
func ReleaseLotteryPerformanceResult(v *LotteryPerformanceResult) {
	v.Rewards = v.Rewards[:0]
	v.ResultCode = 0
	poolLotteryPerformanceResult.Put(v)
}
