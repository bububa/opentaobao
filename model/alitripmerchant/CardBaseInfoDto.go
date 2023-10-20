package alitripmerchant

import (
	"sync"
)

// CardBaseInfoDto 结构体
type CardBaseInfoDto struct {
	// 会员等级
	CardTier string `json:"card_tier,omitempty" xml:"card_tier,omitempty"`
	// 会员卡号
	CardNumber string `json:"card_number,omitempty" xml:"card_number,omitempty"`
	// 会员卡类型
	CardType string `json:"card_type,omitempty" xml:"card_type,omitempty"`
	// 微信会员卡状态
	WechatCardStatus string `json:"wechat_card_status,omitempty" xml:"wechat_card_status,omitempty"`
}

var poolCardBaseInfoDto = sync.Pool{
	New: func() any {
		return new(CardBaseInfoDto)
	},
}

// GetCardBaseInfoDto() 从对象池中获取CardBaseInfoDto
func GetCardBaseInfoDto() *CardBaseInfoDto {
	return poolCardBaseInfoDto.Get().(*CardBaseInfoDto)
}

// ReleaseCardBaseInfoDto 释放CardBaseInfoDto
func ReleaseCardBaseInfoDto(v *CardBaseInfoDto) {
	v.CardTier = ""
	v.CardNumber = ""
	v.CardType = ""
	v.WechatCardStatus = ""
	poolCardBaseInfoDto.Put(v)
}
