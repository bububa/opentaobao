package alitripmerchant

// MemberRightCarouselVo 结构体
type MemberRightCarouselVo struct {
	// 会员权益
	MemberRight []MemberRight `json:"member_right,omitempty" xml:"member_right>member_right,omitempty"`
	// 会员卡英文
	HotelLevelEn string `json:"hotel_level_en,omitempty" xml:"hotel_level_en,omitempty"`
	// 会员卡中文
	HotelLevelZh string `json:"hotel_level_zh,omitempty" xml:"hotel_level_zh,omitempty"`
	// 达到条件
	MemberCardPic string `json:"member_card_pic,omitempty" xml:"member_card_pic,omitempty"`
	// 达到条件
	ReachCondition string `json:"reach_condition,omitempty" xml:"reach_condition,omitempty"`
	// 背景色
	Iconcolor string `json:"iconcolor,omitempty" xml:"iconcolor,omitempty"`
	// 会员卡背景图片
	MemberCardBkPic string `json:"member_card_bk_pic,omitempty" xml:"member_card_bk_pic,omitempty"`
	// 当前等级=false 小于当前等级为空 大于当前等级为true
	IfLockCard bool `json:"if_lock_card,omitempty" xml:"if_lock_card,omitempty"`
}
