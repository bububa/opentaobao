package alitripmerchant

import (
	"sync"
)

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

var poolMemberRightCarouselVo = sync.Pool{
	New: func() any {
		return new(MemberRightCarouselVo)
	},
}

// GetMemberRightCarouselVo() 从对象池中获取MemberRightCarouselVo
func GetMemberRightCarouselVo() *MemberRightCarouselVo {
	return poolMemberRightCarouselVo.Get().(*MemberRightCarouselVo)
}

// ReleaseMemberRightCarouselVo 释放MemberRightCarouselVo
func ReleaseMemberRightCarouselVo(v *MemberRightCarouselVo) {
	v.MemberRight = v.MemberRight[:0]
	v.HotelLevelEn = ""
	v.HotelLevelZh = ""
	v.MemberCardPic = ""
	v.ReachCondition = ""
	v.Iconcolor = ""
	v.MemberCardBkPic = ""
	v.IfLockCard = false
	poolMemberRightCarouselVo.Put(v)
}
