package ascpchannel

import (
	"sync"
)

// Packageitemdtos 结构体
type Packageitemdtos struct {
	// 包裹明细签收信息列表
	PackageItemSignInfoDtoList []Packageitemsigninfodtolist `json:"package_item_sign_info_dto_list,omitempty" xml:"package_item_sign_info_dto_list>packageitemsigninfodtolist,omitempty"`
	// 商品名称
	ItemName string `json:"item_name,omitempty" xml:"item_name,omitempty"`
	// 包裹号
	PackageId int64 `json:"package_id,omitempty" xml:"package_id,omitempty"`
	// 货品id
	ScItemId int64 `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
	// 商品数量
	ItemQuantity int64 `json:"item_quantity,omitempty" xml:"item_quantity,omitempty"`
}

var poolPackageitemdtos = sync.Pool{
	New: func() any {
		return new(Packageitemdtos)
	},
}

// GetPackageitemdtos() 从对象池中获取Packageitemdtos
func GetPackageitemdtos() *Packageitemdtos {
	return poolPackageitemdtos.Get().(*Packageitemdtos)
}

// ReleasePackageitemdtos 释放Packageitemdtos
func ReleasePackageitemdtos(v *Packageitemdtos) {
	v.PackageItemSignInfoDtoList = v.PackageItemSignInfoDtoList[:0]
	v.ItemName = ""
	v.PackageId = 0
	v.ScItemId = 0
	v.ItemQuantity = 0
	poolPackageitemdtos.Put(v)
}
