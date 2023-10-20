package customizemarket

import (
	"sync"
)

// PetProfileDto 结构体
type PetProfileDto struct {
	// 宠物nick
	PetNick string `json:"pet_nick,omitempty" xml:"pet_nick,omitempty"`
	// 宠物种类类型
	PetType string `json:"pet_type,omitempty" xml:"pet_type,omitempty"`
	// 宠物种类名
	PetTypeName string `json:"pet_type_name,omitempty" xml:"pet_type_name,omitempty"`
	// 宠物生日，年-月-日
	PetBirthday string `json:"pet_birthday,omitempty" xml:"pet_birthday,omitempty"`
	// 宠物头像
	PetImage string `json:"pet_image,omitempty" xml:"pet_image,omitempty"`
	// 宠物保险照片
	PetPics string `json:"pet_pics,omitempty" xml:"pet_pics,omitempty"`
	// 宠物档案id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 宠物性别：1男 2 女
	PetSex int64 `json:"pet_sex,omitempty" xml:"pet_sex,omitempty"`
	// 是否绝育：0否，1 是 2 未知
	PetNobaby int64 `json:"pet_nobaby,omitempty" xml:"pet_nobaby,omitempty"`
}

var poolPetProfileDto = sync.Pool{
	New: func() any {
		return new(PetProfileDto)
	},
}

// GetPetProfileDto() 从对象池中获取PetProfileDto
func GetPetProfileDto() *PetProfileDto {
	return poolPetProfileDto.Get().(*PetProfileDto)
}

// ReleasePetProfileDto 释放PetProfileDto
func ReleasePetProfileDto(v *PetProfileDto) {
	v.PetNick = ""
	v.PetType = ""
	v.PetTypeName = ""
	v.PetBirthday = ""
	v.PetImage = ""
	v.PetPics = ""
	v.Id = 0
	v.PetSex = 0
	v.PetNobaby = 0
	poolPetProfileDto.Put(v)
}
