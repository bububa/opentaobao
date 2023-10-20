package tuanhotel

import (
	"sync"
)

// ItemInfoListResult 结构体
type ItemInfoListResult struct {
	// 宝贝信息
	TuanItemOnlines []TuanItemOnlineManagerVo `json:"tuan_item_onlines,omitempty" xml:"tuan_item_onlines>tuan_item_online_manager_vo,omitempty"`
}

var poolItemInfoListResult = sync.Pool{
	New: func() any {
		return new(ItemInfoListResult)
	},
}

// GetItemInfoListResult() 从对象池中获取ItemInfoListResult
func GetItemInfoListResult() *ItemInfoListResult {
	return poolItemInfoListResult.Get().(*ItemInfoListResult)
}

// ReleaseItemInfoListResult 释放ItemInfoListResult
func ReleaseItemInfoListResult(v *ItemInfoListResult) {
	v.TuanItemOnlines = v.TuanItemOnlines[:0]
	poolItemInfoListResult.Put(v)
}
