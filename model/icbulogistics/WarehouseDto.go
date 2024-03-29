package icbulogistics

import (
	"sync"
)

// WarehouseDto 结构体
type WarehouseDto struct {
	// 仓库地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 仓库名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 仓库编码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 联系人
	ContactPerson string `json:"contact_person,omitempty" xml:"contact_person,omitempty"`
	// 联系人电话
	ContactPhone string `json:"contact_phone,omitempty" xml:"contact_phone,omitempty"`
	// 工作时间
	WorkingTime string `json:"working_time,omitempty" xml:"working_time,omitempty"`
	// 邮编
	PostCode string `json:"post_code,omitempty" xml:"post_code,omitempty"`
	// 备注
	Description string `json:"description,omitempty" xml:"description,omitempty"`
}

var poolWarehouseDto = sync.Pool{
	New: func() any {
		return new(WarehouseDto)
	},
}

// GetWarehouseDto() 从对象池中获取WarehouseDto
func GetWarehouseDto() *WarehouseDto {
	return poolWarehouseDto.Get().(*WarehouseDto)
}

// ReleaseWarehouseDto 释放WarehouseDto
func ReleaseWarehouseDto(v *WarehouseDto) {
	v.Address = ""
	v.Name = ""
	v.Code = ""
	v.ContactPerson = ""
	v.ContactPhone = ""
	v.WorkingTime = ""
	v.PostCode = ""
	v.Description = ""
	poolWarehouseDto.Put(v)
}
