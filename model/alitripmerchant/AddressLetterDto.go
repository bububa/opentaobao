package alitripmerchant

import (
	"sync"
)

// AddressLetterDto 结构体
type AddressLetterDto struct {
	// 字母分类列表
	AddressList []AddressSearchDto `json:"address_list,omitempty" xml:"address_list>address_search_dto,omitempty"`
	// 字母
	Letter string `json:"letter,omitempty" xml:"letter,omitempty"`
}

var poolAddressLetterDto = sync.Pool{
	New: func() any {
		return new(AddressLetterDto)
	},
}

// GetAddressLetterDto() 从对象池中获取AddressLetterDto
func GetAddressLetterDto() *AddressLetterDto {
	return poolAddressLetterDto.Get().(*AddressLetterDto)
}

// ReleaseAddressLetterDto 释放AddressLetterDto
func ReleaseAddressLetterDto(v *AddressLetterDto) {
	v.AddressList = v.AddressList[:0]
	v.Letter = ""
	poolAddressLetterDto.Put(v)
}
