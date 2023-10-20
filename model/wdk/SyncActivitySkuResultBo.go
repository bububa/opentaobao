package wdk

import (
	"sync"
)

// SyncActivitySkuResultBo 结构体
type SyncActivitySkuResultBo struct {
	// 商品编码
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// 活动Id
	PromotionId string `json:"promotion_id,omitempty" xml:"promotion_id,omitempty"`
	// 商品池ID
	PoolId int64 `json:"pool_id,omitempty" xml:"pool_id,omitempty"`
	// 版本Id
	VersionId int64 `json:"version_id,omitempty" xml:"version_id,omitempty"`
}

var poolSyncActivitySkuResultBo = sync.Pool{
	New: func() any {
		return new(SyncActivitySkuResultBo)
	},
}

// GetSyncActivitySkuResultBo() 从对象池中获取SyncActivitySkuResultBo
func GetSyncActivitySkuResultBo() *SyncActivitySkuResultBo {
	return poolSyncActivitySkuResultBo.Get().(*SyncActivitySkuResultBo)
}

// ReleaseSyncActivitySkuResultBo 释放SyncActivitySkuResultBo
func ReleaseSyncActivitySkuResultBo(v *SyncActivitySkuResultBo) {
	v.SkuCode = ""
	v.PromotionId = ""
	v.PoolId = 0
	v.VersionId = 0
	poolSyncActivitySkuResultBo.Put(v)
}
