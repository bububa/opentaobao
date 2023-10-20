package alihouse

import (
	"sync"
)

// StoreExtendsInfoDto 结构体
type StoreExtendsInfoDto struct {
	// 外部门店id
	OuterStoreId string `json:"outer_store_id,omitempty" xml:"outer_store_id,omitempty"`
	// 扩展信息
	ExtendsInfo string `json:"extends_info,omitempty" xml:"extends_info,omitempty"`
}

var poolStoreExtendsInfoDto = sync.Pool{
	New: func() any {
		return new(StoreExtendsInfoDto)
	},
}

// GetStoreExtendsInfoDto() 从对象池中获取StoreExtendsInfoDto
func GetStoreExtendsInfoDto() *StoreExtendsInfoDto {
	return poolStoreExtendsInfoDto.Get().(*StoreExtendsInfoDto)
}

// ReleaseStoreExtendsInfoDto 释放StoreExtendsInfoDto
func ReleaseStoreExtendsInfoDto(v *StoreExtendsInfoDto) {
	v.OuterStoreId = ""
	v.ExtendsInfo = ""
	poolStoreExtendsInfoDto.Put(v)
}
