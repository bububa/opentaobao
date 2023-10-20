package drugtrace

import (
	"sync"
)

// EntInfoDto 结构体
type EntInfoDto struct {
	// 企业资质（上传图片）图片建议尺寸：height: 310px;width: 670px;
	EntQualificationPictures []string `json:"ent_qualification_pictures,omitempty" xml:"ent_qualification_pictures>string,omitempty"`
	// 联系方式
	EntContact string `json:"ent_contact,omitempty" xml:"ent_contact,omitempty"`
	// 企业介绍
	EntIntroduction string `json:"ent_introduction,omitempty" xml:"ent_introduction,omitempty"`
}

var poolEntInfoDto = sync.Pool{
	New: func() any {
		return new(EntInfoDto)
	},
}

// GetEntInfoDto() 从对象池中获取EntInfoDto
func GetEntInfoDto() *EntInfoDto {
	return poolEntInfoDto.Get().(*EntInfoDto)
}

// ReleaseEntInfoDto 释放EntInfoDto
func ReleaseEntInfoDto(v *EntInfoDto) {
	v.EntQualificationPictures = v.EntQualificationPictures[:0]
	v.EntContact = ""
	v.EntIntroduction = ""
	poolEntInfoDto.Put(v)
}
