package idleisv

import (
	"sync"
)

// IdleItemBaseApiDo 结构体
type IdleItemBaseApiDo struct {
	// 商品Id（根据此数据进行相应商品下架）
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 服务商Id(已废弃,无需传入)
	ProviderId int64 `json:"provider_id,omitempty" xml:"provider_id,omitempty"`
	// 是否需要sku信息(不需要的业务场景，不要设置为true，会增加查询耗时)
	NeedSku bool `json:"need_sku,omitempty" xml:"need_sku,omitempty"`
}

var poolIdleItemBaseApiDo = sync.Pool{
	New: func() any {
		return new(IdleItemBaseApiDo)
	},
}

// GetIdleItemBaseApiDo() 从对象池中获取IdleItemBaseApiDo
func GetIdleItemBaseApiDo() *IdleItemBaseApiDo {
	return poolIdleItemBaseApiDo.Get().(*IdleItemBaseApiDo)
}

// ReleaseIdleItemBaseApiDo 释放IdleItemBaseApiDo
func ReleaseIdleItemBaseApiDo(v *IdleItemBaseApiDo) {
	v.ItemId = 0
	v.ProviderId = 0
	v.NeedSku = false
	poolIdleItemBaseApiDo.Put(v)
}
