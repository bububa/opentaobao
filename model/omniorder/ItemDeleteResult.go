package omniorder

import (
	"sync"
)

// ItemDeleteResult 结构体
type ItemDeleteResult struct {
	// 重复商品信息
	DuplicateInfos []ItemSkuDuplicateInfo `json:"duplicate_infos,omitempty" xml:"duplicate_infos>item_sku_duplicate_info,omitempty"`
	// itemLightPublishDTO
	ItemLightPublishDTO *ItemLightPublishDto `json:"item_light_publish_d_t_o,omitempty" xml:"item_light_publish_d_t_o,omitempty"`
}

var poolItemDeleteResult = sync.Pool{
	New: func() any {
		return new(ItemDeleteResult)
	},
}

// GetItemDeleteResult() 从对象池中获取ItemDeleteResult
func GetItemDeleteResult() *ItemDeleteResult {
	return poolItemDeleteResult.Get().(*ItemDeleteResult)
}

// ReleaseItemDeleteResult 释放ItemDeleteResult
func ReleaseItemDeleteResult(v *ItemDeleteResult) {
	v.DuplicateInfos = v.DuplicateInfos[:0]
	v.ItemLightPublishDTO = nil
	poolItemDeleteResult.Put(v)
}
