package film

import (
	"sync"
)

// LotteryPerformanceTopParam 结构体
type LotteryPerformanceTopParam struct {
	// 用户对外开放ID
	MixUserId string `json:"mix_user_id,omitempty" xml:"mix_user_id,omitempty"`
	// 活动ID
	SpreadId string `json:"spread_id,omitempty" xml:"spread_id,omitempty"`
	// 幂等ID
	OutBizId string `json:"out_biz_id,omitempty" xml:"out_biz_id,omitempty"`
	// 拓展参数
	BizExtInfo string `json:"biz_ext_info,omitempty" xml:"biz_ext_info,omitempty"`
	// 平台
	Platform int64 `json:"platform,omitempty" xml:"platform,omitempty"`
}

var poolLotteryPerformanceTopParam = sync.Pool{
	New: func() any {
		return new(LotteryPerformanceTopParam)
	},
}

// GetLotteryPerformanceTopParam() 从对象池中获取LotteryPerformanceTopParam
func GetLotteryPerformanceTopParam() *LotteryPerformanceTopParam {
	return poolLotteryPerformanceTopParam.Get().(*LotteryPerformanceTopParam)
}

// ReleaseLotteryPerformanceTopParam 释放LotteryPerformanceTopParam
func ReleaseLotteryPerformanceTopParam(v *LotteryPerformanceTopParam) {
	v.MixUserId = ""
	v.SpreadId = ""
	v.OutBizId = ""
	v.BizExtInfo = ""
	v.Platform = 0
	poolLotteryPerformanceTopParam.Put(v)
}
