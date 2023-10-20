package alsc

import (
	"sync"
)

// StoreKeeperDto 结构体
type StoreKeeperDto struct {
	// 传真
	Fax string `json:"fax,omitempty" xml:"fax,omitempty"`
	// 电话
	Tel string `json:"tel,omitempty" xml:"tel,omitempty"`
	// 移动电话
	Mobile string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 邮编
	ZipCode string `json:"zip_code,omitempty" xml:"zip_code,omitempty"`
	// 门店联系人
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}

var poolStoreKeeperDto = sync.Pool{
	New: func() any {
		return new(StoreKeeperDto)
	},
}

// GetStoreKeeperDto() 从对象池中获取StoreKeeperDto
func GetStoreKeeperDto() *StoreKeeperDto {
	return poolStoreKeeperDto.Get().(*StoreKeeperDto)
}

// ReleaseStoreKeeperDto 释放StoreKeeperDto
func ReleaseStoreKeeperDto(v *StoreKeeperDto) {
	v.Fax = ""
	v.Tel = ""
	v.Mobile = ""
	v.ZipCode = ""
	v.Name = ""
	poolStoreKeeperDto.Put(v)
}
