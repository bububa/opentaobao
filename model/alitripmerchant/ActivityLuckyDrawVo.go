package alitripmerchant

import (
	"sync"
)

// ActivityLuckyDrawVo 结构体
type ActivityLuckyDrawVo struct {
	// 奖品
	ActivityGoodsList []ActivityGoodsList `json:"activity_goods_list,omitempty" xml:"activity_goods_list>activity_goods_list,omitempty"`
	// 老虎机奖品展示对象
	PrizeDisplays []PrizeDisplay `json:"prize_displays,omitempty" xml:"prize_displays>prize_display,omitempty"`
	// 未中奖图片
	UnWinningImages string `json:"un_winning_images,omitempty" xml:"un_winning_images,omitempty"`
}

var poolActivityLuckyDrawVo = sync.Pool{
	New: func() any {
		return new(ActivityLuckyDrawVo)
	},
}

// GetActivityLuckyDrawVo() 从对象池中获取ActivityLuckyDrawVo
func GetActivityLuckyDrawVo() *ActivityLuckyDrawVo {
	return poolActivityLuckyDrawVo.Get().(*ActivityLuckyDrawVo)
}

// ReleaseActivityLuckyDrawVo 释放ActivityLuckyDrawVo
func ReleaseActivityLuckyDrawVo(v *ActivityLuckyDrawVo) {
	v.ActivityGoodsList = v.ActivityGoodsList[:0]
	v.PrizeDisplays = v.PrizeDisplays[:0]
	v.UnWinningImages = ""
	poolActivityLuckyDrawVo.Put(v)
}
