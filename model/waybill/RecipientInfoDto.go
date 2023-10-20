package waybill

import (
	"sync"
)

// RecipientInfoDto 结构体
type RecipientInfoDto struct {
	// 手机号码
	Mobile string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 姓名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 固定电话
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
	// 淘宝订单收件人ID (Open Addressee ID)，长度不超过128个字符，淘宝订单加密情况用于解密。
	Oaid string `json:"oaid,omitempty" xml:"oaid,omitempty"`
	// 电商平台真实交易订单号，针对电商平台订单隐私加密场景使用，非必填，如果填写则必须是电商平台真实的交易订单ID
	Tid string `json:"tid,omitempty" xml:"tid,omitempty"`
	// 菜鸟解密地址ID，用于电商平台收件人信息加密的场景使用，非订单加密场景请勿使用。
	Caid string `json:"caid,omitempty" xml:"caid,omitempty"`
	// 地址
	Address *AddressDto `json:"address,omitempty" xml:"address,omitempty"`
}

var poolRecipientInfoDto = sync.Pool{
	New: func() any {
		return new(RecipientInfoDto)
	},
}

// GetRecipientInfoDto() 从对象池中获取RecipientInfoDto
func GetRecipientInfoDto() *RecipientInfoDto {
	return poolRecipientInfoDto.Get().(*RecipientInfoDto)
}

// ReleaseRecipientInfoDto 释放RecipientInfoDto
func ReleaseRecipientInfoDto(v *RecipientInfoDto) {
	v.Mobile = ""
	v.Name = ""
	v.Phone = ""
	v.Oaid = ""
	v.Tid = ""
	v.Caid = ""
	v.Address = nil
	poolRecipientInfoDto.Put(v)
}
