package alitripmerchant

// MemberCardDto 结构体
type MemberCardDto struct {
	// 会员权益
	MemberRights []MemberRights `json:"member_rights,omitempty" xml:"member_rights>member_rights,omitempty"`
	// 飞猪等级
	FliggyLevel string `json:"fliggy_level,omitempty" xml:"fliggy_level,omitempty"`
	// 会员卡图片地址
	MemberCardPic string `json:"member_card_pic,omitempty" xml:"member_card_pic,omitempty"`
	// 酒店等级
	HotelLevel string `json:"hotel_level,omitempty" xml:"hotel_level,omitempty"`
	// 权益描述
	CardName string `json:"card_name,omitempty" xml:"card_name,omitempty"`
	// 会员卡ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}
