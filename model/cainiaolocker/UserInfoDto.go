package cainiaolocker

import (
	"sync"
)

// UserInfoDto 结构体
type UserInfoDto struct {
	// 手机号码（手机号和固定电话不能同时为空），长度小于20
	Mobile string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 姓名，长度小于40
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 固定电话（手机号和固定电话不能同时为空），长度小于20
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
	// 菜鸟解密地址ID，用于电商平台收件人信息加密的场景使用，非订单加密场景请勿使用。
	Caid string `json:"caid,omitempty" xml:"caid,omitempty"`
	// 淘宝订单收件人ID (Open Addressee ID)，长度不超过128个字符，淘宝订单加密情况用于解密。
	Oaid string `json:"oaid,omitempty" xml:"oaid,omitempty"`
	// 电商平台真实交易订单号，针对电商平台订单隐私加密场景使用，非必填，如果填写则必须是电商平台真实的交易订单ID
	Tid string `json:"tid,omitempty" xml:"tid,omitempty"`
	// 发货地址需要通过&lt;a href=&#34;http://open.taobao.com/doc2/detail.htm?spm=a219a.7629140.0.0.3OFCPk&amp;treeId=17&amp;articleId=104860&amp;docType=1&#34;&gt;search接口&lt;/a&gt;
	Address *AddressDto `json:"address,omitempty" xml:"address,omitempty"`
}

var poolUserInfoDto = sync.Pool{
	New: func() any {
		return new(UserInfoDto)
	},
}

// GetUserInfoDto() 从对象池中获取UserInfoDto
func GetUserInfoDto() *UserInfoDto {
	return poolUserInfoDto.Get().(*UserInfoDto)
}

// ReleaseUserInfoDto 释放UserInfoDto
func ReleaseUserInfoDto(v *UserInfoDto) {
	v.Mobile = ""
	v.Name = ""
	v.Phone = ""
	v.Caid = ""
	v.Oaid = ""
	v.Tid = ""
	v.Address = nil
	poolUserInfoDto.Put(v)
}
