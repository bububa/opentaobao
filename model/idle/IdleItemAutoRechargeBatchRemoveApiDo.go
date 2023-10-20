package idle

import (
	"sync"
)

// IdleItemAutoRechargeBatchRemoveApiDo 结构体
type IdleItemAutoRechargeBatchRemoveApiDo struct {
	// 商品id列表
	ItemIds []int64 `json:"item_ids,omitempty" xml:"item_ids>int64,omitempty"`
}

var poolIdleItemAutoRechargeBatchRemoveApiDo = sync.Pool{
	New: func() any {
		return new(IdleItemAutoRechargeBatchRemoveApiDo)
	},
}

// GetIdleItemAutoRechargeBatchRemoveApiDo() 从对象池中获取IdleItemAutoRechargeBatchRemoveApiDo
func GetIdleItemAutoRechargeBatchRemoveApiDo() *IdleItemAutoRechargeBatchRemoveApiDo {
	return poolIdleItemAutoRechargeBatchRemoveApiDo.Get().(*IdleItemAutoRechargeBatchRemoveApiDo)
}

// ReleaseIdleItemAutoRechargeBatchRemoveApiDo 释放IdleItemAutoRechargeBatchRemoveApiDo
func ReleaseIdleItemAutoRechargeBatchRemoveApiDo(v *IdleItemAutoRechargeBatchRemoveApiDo) {
	v.ItemIds = v.ItemIds[:0]
	poolIdleItemAutoRechargeBatchRemoveApiDo.Put(v)
}
