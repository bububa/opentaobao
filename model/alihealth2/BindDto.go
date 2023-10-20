package alihealth2

import (
	"sync"
)

// BindDto 结构体
type BindDto struct {
	// sp(服务商)门店ID
	SpStoreId string `json:"sp_store_id,omitempty" xml:"sp_store_id,omitempty"`
	// sp(服务商)商品ID
	SpItemId string `json:"sp_item_id,omitempty" xml:"sp_item_id,omitempty"`
	// 天猫商品ID
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 天猫门店ID
	StoreId int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
}

var poolBindDto = sync.Pool{
	New: func() any {
		return new(BindDto)
	},
}

// GetBindDto() 从对象池中获取BindDto
func GetBindDto() *BindDto {
	return poolBindDto.Get().(*BindDto)
}

// ReleaseBindDto 释放BindDto
func ReleaseBindDto(v *BindDto) {
	v.SpStoreId = ""
	v.SpItemId = ""
	v.ItemId = 0
	v.StoreId = 0
	poolBindDto.Put(v)
}
