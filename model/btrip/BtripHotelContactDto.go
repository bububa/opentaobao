package btrip

import (
	"sync"
)

// BtripHotelContactDto 结构体
type BtripHotelContactDto struct {
	// 邮箱
	Email string `json:"email,omitempty" xml:"email,omitempty"`
	// 入住人姓名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 入住人电话
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
}

var poolBtripHotelContactDto = sync.Pool{
	New: func() any {
		return new(BtripHotelContactDto)
	},
}

// GetBtripHotelContactDto() 从对象池中获取BtripHotelContactDto
func GetBtripHotelContactDto() *BtripHotelContactDto {
	return poolBtripHotelContactDto.Get().(*BtripHotelContactDto)
}

// ReleaseBtripHotelContactDto 释放BtripHotelContactDto
func ReleaseBtripHotelContactDto(v *BtripHotelContactDto) {
	v.Email = ""
	v.Name = ""
	v.Phone = ""
	poolBtripHotelContactDto.Put(v)
}
