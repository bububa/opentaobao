package tbk

import (
	"sync"
)

// RecommendItemList 结构体
type RecommendItemList struct {
	// 商品链接
	Url string `json:"url,omitempty" xml:"url,omitempty"`
	// 权益推荐商品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
}

var poolRecommendItemList = sync.Pool{
	New: func() any {
		return new(RecommendItemList)
	},
}

// GetRecommendItemList() 从对象池中获取RecommendItemList
func GetRecommendItemList() *RecommendItemList {
	return poolRecommendItemList.Get().(*RecommendItemList)
}

// ReleaseRecommendItemList 释放RecommendItemList
func ReleaseRecommendItemList(v *RecommendItemList) {
	v.Url = ""
	v.ItemId = 0
	poolRecommendItemList.Put(v)
}
