package alitripmerchant

// MemberCardDetailVo 结构体
type MemberCardDetailVo struct {
	// 会员权益集合
	MemberRightList []MemberRightVo `json:"member_right_list,omitempty" xml:"member_right_list>member_right_vo,omitempty"`
	// 卡图片
	MemberCardPic string `json:"member_card_pic,omitempty" xml:"member_card_pic,omitempty"`
	// 英文等级
	HotelLevelEn string `json:"hotel_level_en,omitempty" xml:"hotel_level_en,omitempty"`
	// 会员升级条件
	ReachCondition string `json:"reach_condition,omitempty" xml:"reach_condition,omitempty"`
	// 图标颜色
	IconColor string `json:"icon_color,omitempty" xml:"icon_color,omitempty"`
	// 会员卡描述
	CardDesc string `json:"card_desc,omitempty" xml:"card_desc,omitempty"`
	// 背景图片
	BackGroundPic string `json:"back_ground_pic,omitempty" xml:"back_ground_pic,omitempty"`
	// 会员等级中文名称
	HotelLevelZh string `json:"hotel_level_zh,omitempty" xml:"hotel_level_zh,omitempty"`
	// 个人中心背景图片
	PersonalCenterPic string `json:"personal_center_pic,omitempty" xml:"personal_center_pic,omitempty"`
}
