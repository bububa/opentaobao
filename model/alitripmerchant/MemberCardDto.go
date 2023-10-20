package alitripmerchant

import (
	"sync"
)

// MemberCardDto 结构体
type MemberCardDto struct {
	// 会员权益
	MemberRights []MemberRights `json:"member_rights,omitempty" xml:"member_rights>member_rights,omitempty"`
	// 酒店等级
	HotelLevel string `json:"hotel_level,omitempty" xml:"hotel_level,omitempty"`
	// 飞猪等级
	FliggyLevel string `json:"fliggy_level,omitempty" xml:"fliggy_level,omitempty"`
	// 会员卡图片地址
	MemberCardPic string `json:"member_card_pic,omitempty" xml:"member_card_pic,omitempty"`
	// 权益描述
	CardName string `json:"card_name,omitempty" xml:"card_name,omitempty"`
	// 添加会员卡JSOS入参
	CardExt string `json:"card_ext,omitempty" xml:"card_ext,omitempty"`
	// 加密code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 模板id
	CardId string `json:"card_id,omitempty" xml:"card_id,omitempty"`
	// 会员卡ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolMemberCardDto = sync.Pool{
	New: func() any {
		return new(MemberCardDto)
	},
}

// GetMemberCardDto() 从对象池中获取MemberCardDto
func GetMemberCardDto() *MemberCardDto {
	return poolMemberCardDto.Get().(*MemberCardDto)
}

// ReleaseMemberCardDto 释放MemberCardDto
func ReleaseMemberCardDto(v *MemberCardDto) {
	v.MemberRights = v.MemberRights[:0]
	v.HotelLevel = ""
	v.FliggyLevel = ""
	v.MemberCardPic = ""
	v.CardName = ""
	v.CardExt = ""
	v.Code = ""
	v.CardId = ""
	v.Id = 0
	poolMemberCardDto.Put(v)
}
