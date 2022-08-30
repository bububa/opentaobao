package alitripmerchant

// ActivityLuckyDrawVO 结构体
type ActivityLuckyDrawVO struct {
	// 奖品
	ActivityGoodsList []ActivityGoodsList `json:"activity_goods_list,omitempty" xml:"activity_goods_list>activity_goods_list,omitempty"`
	// 老虎机奖品展示对象
	PrizeDisplays []PrizeDisplay `json:"prize_displays,omitempty" xml:"prize_displays>prize_display,omitempty"`
	// 未中奖图片
	UnWinningImages string `json:"un_winning_images,omitempty" xml:"un_winning_images,omitempty"`
}
