package security

import (
	"sync"
)

// RpidCardBo 结构体
type RpidCardBo struct {
	// 类型
	CardType string `json:"card_type,omitempty" xml:"card_type,omitempty"`
	// 身份证背面照的URL地址
	UrlBackImage string `json:"url_back_image,omitempty" xml:"url_back_image,omitempty"`
	// 身份证正面照的URL地址
	UrlFrontImage string `json:"url_front_image,omitempty" xml:"url_front_image,omitempty"`
	// 身份证上生日
	BirthDay string `json:"birth_day,omitempty" xml:"birth_day,omitempty"`
	// 身份证住址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 身份证上的名字
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 公民身份证号码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 性别
	Sex *RpSex `json:"sex,omitempty" xml:"sex,omitempty"`
	// RPIDCardImage
	RPIDCardImage *RpidCardImage `json:"r_p_i_d_card_image,omitempty" xml:"r_p_i_d_card_image,omitempty"`
}

var poolRpidCardBo = sync.Pool{
	New: func() any {
		return new(RpidCardBo)
	},
}

// GetRpidCardBo() 从对象池中获取RpidCardBo
func GetRpidCardBo() *RpidCardBo {
	return poolRpidCardBo.Get().(*RpidCardBo)
}

// ReleaseRpidCardBo 释放RpidCardBo
func ReleaseRpidCardBo(v *RpidCardBo) {
	v.CardType = ""
	v.UrlBackImage = ""
	v.UrlFrontImage = ""
	v.BirthDay = ""
	v.Address = ""
	v.Name = ""
	v.Code = ""
	v.Sex = nil
	v.RPIDCardImage = nil
	poolRpidCardBo.Put(v)
}
