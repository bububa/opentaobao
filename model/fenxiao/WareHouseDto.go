package fenxiao

import (
	"sync"
)

// WareHouseDto 结构体
type WareHouseDto struct {
	// 详细地址描述
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 仓库地址信息,格式 :省~市~区县
	AddressAreaName string `json:"address_area_name,omitempty" xml:"address_area_name,omitempty"`
	// 别名
	AliasName string `json:"alias_name,omitempty" xml:"alias_name,omitempty"`
	// 联系人
	Contact string `json:"contact,omitempty" xml:"contact,omitempty"`
	// 电话号码
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
	// 邮编
	PostCode string `json:"post_code,omitempty" xml:"post_code,omitempty"`
	// 仓库编码，仅允许使用字母、数字、下划线，并且在6-30个字符内
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
	// 仓库名称
	StoreName string `json:"store_name,omitempty" xml:"store_name,omitempty"`
	// 操作类型，新增:ADD;修改:UPDATE
	OperateType string `json:"operate_type,omitempty" xml:"operate_type,omitempty"`
}

var poolWareHouseDto = sync.Pool{
	New: func() any {
		return new(WareHouseDto)
	},
}

// GetWareHouseDto() 从对象池中获取WareHouseDto
func GetWareHouseDto() *WareHouseDto {
	return poolWareHouseDto.Get().(*WareHouseDto)
}

// ReleaseWareHouseDto 释放WareHouseDto
func ReleaseWareHouseDto(v *WareHouseDto) {
	v.Address = ""
	v.AddressAreaName = ""
	v.AliasName = ""
	v.Contact = ""
	v.Phone = ""
	v.PostCode = ""
	v.StoreCode = ""
	v.StoreName = ""
	v.OperateType = ""
	poolWareHouseDto.Put(v)
}
